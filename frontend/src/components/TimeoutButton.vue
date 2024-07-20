<template>
    <button :class="{'button-inactive': !timeoutActive, 'button-active': timeoutActive}" @click="buttonClicked">
        <CopyActiveIcon class="button-timer" v-if="timeoutActive" :count-down="props.countdownDuration"></CopyActiveIcon>
        <span class="button-text" v-else>{{  props.name }}</span>
    </button>
</template>

<script setup lang="ts">
import { ref, toRefs } from 'vue'
import CopyActiveIcon from './CopyActiveIcon.vue';

const timeoutActive = ref(false)

const props = defineProps<{
   name: string
   activeColor: string
   countdownDuration: number 
}>()

const { activeColor } = toRefs(props)


const emit = defineEmits<{
    "countdownStarted": []
    "countdownEnded": []
}>()

const buttonClicked = () => {
    emit("countdownStarted")
    timeoutActive.value = true
    const intervalId = setTimeout(() => {
        emit("countdownEnded")
        timeoutActive.value = false 
    }, props.countdownDuration * 1000)
}

</script>

<style scoped>
button {
    display: flex;
    text-align: center;
    border-radius: 5px;
    flex-grow: 0;
    background-color: white;
    color: #555;
    border: 1px solid #CCC;
    font-size: 2em;
    padding: 10px 15px;
    margin-bottom: 5px;
}

.button-inactive {

}

.button-inactive:hover {
    cursor: pointer;
    background-color: #555;
    color: white;
}

.button-active {
    color: "white";
    background-color: v-bind(activeColor);
    /* rgb(7, 166, 7); */
}


.button-inactive:hover {
    cursor: pointer;
    background-color: #555;
    color: white;
}

.button-text {
    margin-right: auto;
    margin-left: auto;
}

.button-timer {
    margin-left: auto;
    margin-right: auto;
}
</style>