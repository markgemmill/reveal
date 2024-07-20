<template>
    <GearIcon class="settings-button" @click="openSettings"></GearIcon>
    <DockIcon class="dock-button" @click="moveToDockPosition"></DockIcon>
    <ResetIcon class="reset-button" @click="resetPassword"></ResetIcon>
    <div v-if="!active" class="page page-start">
        <button class="start-button" @click="startPassword">START</button>
    </div>
    <div v-else class="page page-password">
        <div class="panel-password">
            <PasswordDisplay 
              :password="password" 
              :segmentSize="segmentSize" 
              :visible="visible"
              ></PasswordDisplay>
        </div>
        <div class="panel-buttons">
            <TimeoutButton 
              :name="'Copy'" 
              :active-color="'#80f41a'"
              :countdown-duration="clipboardLifetime"
              @countdown-started="setPasswordToClipboard" 
              @countdown-ended="removePasswordFromClipboard"
              ></TimeoutButton>
            <TimeoutButton 
              :name="'Reveal'" 
              :active-color="'#951c04'"
              :countdown-duration="displayLifetime"
              @countdown-started="displayPassword" 
              @countdown-ended="hidePassword"
              ></TimeoutButton>
        </div>
    </div>
    <SettingsPage 
      v-if="displaySettings"
      v-model:clipboard="clipboardLifetime"
      v-model:display="displayLifetime"
      v-model:password="segmentSize"
      v-model:version="version"
      v-model:dockPosition="dockPosition"
      @close="closeSettings"
      ></SettingsPage>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import PasswordDisplay from './PasswordDisplay.vue'
import * as App from '../../wailsjs/go/main/App'
import * as runtime from '../../wailsjs/runtime'
import TimeoutButton from "./TimeoutButton.vue"
import GearIcon from "./GearIcon.vue"
import ResetIcon from "./ResetIcon.vue"
import DockIcon from './DockIcon.vue'
import SettingsPage from "./SettingsPage.vue" 
import { DockPosition } from './types'

const password = ref("")
const active = ref(false)
const visible = ref(false)
const valueOnClipboard = ref(false)

const isLoading = ref(true)
const displaySettings = ref(false)
const clipboardLifetime = ref(20)
const displayLifetime = ref(20) 
const segmentSize = ref(5)
const version = ref("")
const dockPosition = ref<DockPosition>("center")
const position = ref<{x: number, y: number}>({x: 0, y: 0})
const size = ref<{h: number, w: number}>({h: 0, w: 0})
const mountCompleted = ref(0)
let passwordDate = new Date()

const checkDateChange = () => {
    const currentDate = new Date()
    if (
        currentDate.getFullYear() !== passwordDate.getFullYear() || 
        currentDate.getMonth() !== passwordDate.getMonth() ||
        currentDate.getDate() !== passwordDate.getDate()
    ) {
        passwordDate = currentDate
        resetPassword()
    } 
}

const resetPassword = () => {
    password.value = ""
    active.value = false
}

const moveToDockPosition = () => {
   dock() 
}

const openSettings = () => {
    displaySettings.value = true;
}

const closeSettings = () => {
    displaySettings.value = false
}

const startPassword = () => {
   // copy password from clipboard 
   App.GetClipboard().then((clipboardValue: string) => {
        password.value = clipboardValue 
        active.value = true
        App.ClearClipboard()
        displayPassword()
        console.debug("set start timeout... " + displayLifetime.value)
        setTimeout(() => {
            hidePassword() 
        }, displayLifetime.value * 1000)
   }) 
}

const displayPassword = () => {
    visible.value = true
}

const hidePassword = () => {
    visible.value = false 
}

const setPasswordToClipboard = () => {
    App.PutClipboard(password.value).then(() => {
        console.debug("password copied to clipboard!") 
        valueOnClipboard.value = true
    }) 
}

const removePasswordFromClipboard = () => {
    App.ClearClipboard()
    valueOnClipboard.value = false 
}

const ensureInt = (value: number | string) => {
    return Number.parseInt(value.toString())
}

watch([
    clipboardLifetime,
    displayLifetime,
    segmentSize,
    dockPosition
], () => {
    if (isLoading.value == true) return
    console.debug("saving config changes...")

    App.WriteConfig(
        ensureInt(displayLifetime.value), 
        ensureInt(clipboardLifetime.value), 
        ensureInt(segmentSize.value),
        dockPosition.value
    ) 
})

watch(dockPosition, (value) => {
    console.debug(`dock position changed to: ${value}`) 
})

watch(mountCompleted, (value) => {
    runtime.LogDebug(`mountCompleted: ${value}`)
    if (value >= 3) {
        runtime.LogDebug(`mountCompleted calling dock()!`)
        dock()
        isLoading.value = false 
    } 
})

const dock = () => {
    let x = 0
    let y = 0
    runtime.LogDebug(`Current dock position: ${dockPosition.value}`)
    runtime.LogDebug(`x: ${x}, y: ${y}`)
    switch (dockPosition.value) {
        case "center":
            x = (screen.width / 2) - (size.value.w / 2)
            y = (screen.height / 2) - (size.value.h / 2) 
            break;
        case "ul":
            x = 0 
            y = 0
            break;
        case "ur":
            x = screen.width - size.value.w
            y = 0
            break;
        case "ll":
            x = 0
            y = screen.height - size.value.h 
            break;
        case "lr":
            x = screen.width - size.value.w
            y = screen.height - size.value.h 
            break;
    }
    console.debug(`setting screen position to: ${dockPosition.value} (${x}, ${y})`)
    runtime.LogDebug(`setting screen position to: ${dockPosition.value} (${x}, ${y})`)
    runtime.WindowSetPosition(x, y)
}

onMounted(() => {
    mountCompleted.value = 0
    setInterval(() => {
        checkDateChange()
    }, 3600000)

    App.GetTimeout().then((value: number) => {
        console.debug(`Set timeout to: ${value}`)
        displayLifetime.value = value 
    })
    App.GetCopyTimeout().then((value: number) => {
        console.debug(`Set copy timeout to: ${value}`)
        clipboardLifetime.value = value
    })
    App.GetSegmentSize().then((value: number) => {
        segmentSize.value = value 
    })
    App.GetVersion().then((value: string) => {
        version.value = `v${value}`        
    })
    App.GetDockPosition().then((value: string) => {
        console.debug(`set dock position to: ${value}`)
        dockPosition.value = value as DockPosition 
    })

    let screen = null 
    runtime.ScreenGetAll().then((value) => {
        screen = value[0]
        console.debug(value)
        console.debug(`width:  ${screen.width}`)
        console.debug(`height: ${screen.height}`)
        mountCompleted.value += 1
    })

    runtime.WindowGetPosition().then((value) => {
        console.debug(value) 
        position.value.x = value.x
        position.value.y = value.y
        mountCompleted.value += 1
    })

    runtime.WindowGetSize().then((value) => {
        console.debug(value) 
        size.value.h = value.h
        size.value.w = value.w
        mountCompleted.value += 1
    })

})

</script>

<style scoped>
</style>