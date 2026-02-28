<script lang="ts" setup>
import { ref } from 'vue'

const code = ref(`# 数据清洗模板
library(dplyr)
library(tidyr)

# 处理缺失值
data <- data %>%
  drop_na() %>%
  filter(complete.cases(.))

# 移除异常值
data <- data %>%
  mutate(across(where(is.numeric), ~ {
    qnt <- quantile(., probs = c(0.25, 0.75), na.rm = TRUE)
    H <- 1.5 * IQR(., na.rm = TRUE)
    .[. < (qnt[1] - H) | . > (qnt[2] + H)] <- NA
    .
  }))
`)
</script>

<template>
  <div class="templates-view h-100 d-flex flex-column">
    <!-- 代码编辑器 -->
    <div class="editor-header pa-3 d-flex align-center" style="background: rgb(var(--v-theme-surface)); border-bottom: 1px solid rgb(var(--v-theme-border));">
      <v-icon icon="mdi-language-r" color="primary" class="mr-2" />
      <span class="text-subtitle-2">数据清洗.R</span>
      <v-spacer />
      <v-btn size="small" variant="outlined" class="mr-2">
        <v-icon start>mdi-play</v-icon>
        运行
      </v-btn>
      <v-btn size="small" variant="outlined">
        <v-icon start>mdi-open-in-new</v-icon>
        外部IDE
      </v-btn>
    </div>
    
    <div class="code-editor flex-grow-1" style="background: rgb(var(--v-theme-background)); padding: 16px;">
      <pre style="font-family: 'JetBrains Mono', monospace; font-size: 14px; line-height: 1.6; color: rgb(var(--v-theme-on-surface));">{{ code }}</pre>
    </div>
    
    <!-- 终端输出区域 -->
    <div class="terminal-output pa-3" style="background: rgb(var(--v-theme-surface-variant)); border-top: 1px solid rgb(var(--v-theme-border)); height: 200px;">
      <div class="d-flex align-center mb-2">
        <v-icon icon="mdi-console" size="small" color="secondary" class="mr-2" />
        <span class="text-caption">终端输出</span>
        <v-spacer />
        <v-btn size="x-small" variant="text">清空</v-btn>
      </div>
      <div class="output-content" style="font-family: 'JetBrains Mono', monospace; font-size: 12px; color: rgb(var(--v-theme-on-surface)); opacity: 0.6;">
        点击"运行"按钮执行代码...
      </div>
    </div>
  </div>
</template>

<style scoped>
.code-editor pre {
  margin: 0;
  white-space: pre-wrap;
}
</style>
