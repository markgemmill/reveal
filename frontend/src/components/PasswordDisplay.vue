<template>
    <div class="password-container" :style="styles">
        <div v-if="props.visible" v-for="letter, index in password" class="password-character">{{ letter }}</div>
        <div v-else v-for="letter, index in hiddenPassword" class="password-character">{{ letter }}</div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'

const password = ref("")
const hiddenPassword = ref("")

const props = defineProps<{
    password: string
    segmentSize: number
    visible: boolean
}>()

const styles = reactive({
    gridTemplateColumns: "auto ".repeat(props.segmentSize) 
})

onMounted(() => {
    password.value = props.password
    hiddenPassword.value = "*".repeat(props.password.length)
})

</script>

<style scoped>
</style>