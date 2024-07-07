<template>
    <div v-if="!active" class="page page-start">
        <button class="start-button" @click="setPassword">START</button>
    </div>
    <div v-else class="page page-password">
        <div class="panel-password">
            <PasswordDisplay class="password" :password="password" :visible="visible"></PasswordDisplay>
        </div>
        <div class="panel-buttons">
            <button class="copy-button" @click="copyPassword">Copy</button>
            <button class="reveal-button" @click="displayPassword" :disabled="visible">Reveal</button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted} from 'vue'
import PasswordDisplay from './PasswordDisplay.vue'
import { GetClipboard, PutClipboard } from '../../wailsjs/go/main/App'

const password = ref("")
const active = ref(false)
const visible = ref(false)
const timeout = 6000
let passwordDate = new Date()

const checkDateChange = (date: Date) => {
    const currentDate = new Date()
    if (
        currentDate.getFullYear() !== date.getFullYear() || 
        currentDate.getMonth() !== date.getMonth() ||
        currentDate.getDate() !== date.getDate()
    ) {
        passwordDate = currentDate
        resetPassword()
    } 
}

const resetPassword = () => {
    password.value = ""
    active.value = false
}

const setPassword = () => {
   // copy password from clipboard 
   GetClipboard().then((clipboardValue: string) => {
    password.value = clipboardValue 
    active.value = true
    displayPassword()
   }) 
}

const displayPassword = () => {
   visible.value = true
   setTimeout(() => {
        visible.value = false
   }, timeout)
}

const copyPassword = () => {
    PutClipboard(password.value).then(() => {
        console.debug("password copied to clipboard!") 
    }) 
}

onMounted(() => {
    setInterval(() => {
        checkDateChange(new Date())
    }, 3600000)
})

</script>

<style scoped>
</style>