<template>
    <div class="password-container" :style="styles">
        <div v-if="visible" v-for="letter, index in password" class="password-character">{{ letter }}</div>
        <div v-else v-for="letter, index in hiddenPassword" class="password-character">{{ letter }}</div>
    </div>
</template>

<script setup lang="ts">
import { ref, toRefs, reactive, computed, onMounted } from 'vue'

const password = ref("")
const hiddenPassword = ref("")

const props = defineProps<{
    password: string
    segmentSize: number
    visible: boolean
}>()

const { visible, segmentSize } = toRefs(props)

const styles = computed(() => {
    return {
        gridTemplateColumns: "auto ".repeat(segmentSize.value) 
    }
})

onMounted(() => {
    password.value = props.password
    hiddenPassword.value = "*".repeat(props.password.length)
})

</script>

<style scoped>
</style>