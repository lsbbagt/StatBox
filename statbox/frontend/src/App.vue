<script lang="ts" setup>
import { ref } from 'vue'

const message = ref('StatBox 已成功启动！')
const count = ref(0)
</script>

<template>
  <v-app>
    <v-main>
      <v-container class="fill-height">
        <v-row justify="center" align="center">
          <v-col cols="12" sm="8" md="6">
            <v-card class="text-center pa-8" elevation="0">
              <v-icon icon="mdi-chart-box" size="80" color="primary" class="mb-4" />
              <h1 class="text-h3 mb-4">StatBox</h1>
              <p class="text-body-1 mb-6">{{ message }}</p>
              
              <v-btn
                color="primary"
                size="large"
                @click="count++"
              >
                点击计数: {{ count }}
              </v-btn>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>
html {
  overflow-y: hidden !important;
}
</style>
