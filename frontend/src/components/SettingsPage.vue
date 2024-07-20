<template>
<div class="settings-page">
    <div class="settings-form">
        <div class="close-button">
            <svg 
              xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-x-lg" viewBox="0 0 16 16"
              @click="close"
              >
                <path d="M2.146 2.854a.5.5 0 1 1 .708-.708L8 7.293l5.146-5.147a.5.5 0 0 1 .708.708L8.707 8l5.147 5.146a.5.5 0 0 1-.708.708L8 8.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 8z"/>
            </svg>
        </div>
        <div class="setting">
            <div class="slider">
                <label for="displayDuration">Display Duration</label>
                <input 
                  v-model="displayDuration" 
                  type="range" 
                  id="displayDuration" 
                  min=10 
                  max=120>
            </div>
            <div class="value">
                {{ displayDuration }} 
            </div>
        </div>
        <div class="setting">
            <div class="slider">
                <label for="clipboardLifetime">Clipboard Life</label>
                <input 
                  v-model="clipboardLifetime" 
                  type="range" 
                  id="clipboardLifetime" 
                  min=5 
                  max=60>
            </div>
            <div class="value">
              {{  clipboardLifetime }}
            </div>
        </div>
        <div class="setting">
            <div class="slider">
                <label for="passwordSplitSize">Password Split Size</label>
                <input 
                  v-model="passwordSplitSize" 
                  type="range" 
                  id="passwordSplitSize" 
                  min=3 
                  max=5>
            </div>
            <div class="value">
               {{ passwordSplitSize }} 
            </div>
        </div>
        <div class="setting">
            <div class="slider">
                <label for="passwordSplitSize">Dock Position</label>
                <DockPositions v-model:position="dockPosition"></DockPositions>
            </div>
        </div>
        <div class="version">{{ version }}</div>
    </div>
</div>    
</template>

<script setup lang="ts">
import DockPositions from "./DockPositions.vue"
import type { DockPosition } from "./types"
const displayDuration = defineModel<number>('display', {required: true})
const clipboardLifetime = defineModel<number>('clipboard', {required: true}) 
const passwordSplitSize = defineModel<number>('password', {required: true})
const dockPosition = defineModel<DockPosition>('dockPosition', {required: true})
const version = defineModel<string>('version', {required: true})

const emit = defineEmits<{
    "close": []
}>()

const close = () => {
    emit("close")
}

</script>

<style scoped>
.settings-page {
    z-index: 10000;
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: white;
}
.settings-form {
    width: 80%;
    /* height: 80%; */
    margin-left: auto;
    margin-right: auto;
    margin-top: 20px;
    padding: 10px;
    padding-top: 50px;
    border: 1px solid #CCC;
    border-radius: 5px;
    background-color: white;
    display: flex;
    flex-direction: column;
    padding: 20px;
}
.setting {
    display: flex;
    flex-direction: row;
    align-items: start;
    margin-bottom: 20px;
}

.slider {
    display: flex;
    flex-grow: 1;
    flex-direction: column;
    align-items: start;
}

.slider label {
    font-weight: bold;
    margin-bottom: 10px;
}
.value {
    flex-grow: 0;
    height: 100%;
    font-size: 2em;
    margin-top: auto;
}
.close-button {
    display: flex;
    flex-direction: row;
    align-self: end;
    margin-bottom: 20px;
}
.version {
    margin-top: 1em;
    color: #AAA;
}

</style>