<script lang="ts" setup>
import { ref } from 'vue'

const commandContent = ref(`description: 使用最优参数编译LaTeX文档

pdflatex -interaction=nonstopmode -shell-escape main.tex
biber main
pdflatex main.tex
pdflatex main.tex
`)
</script>

<template>
  <div class="commands-view h-100">
    <v-card class="ma-4" elevation="0">
      <v-card-title class="d-flex align-center pb-2">
        <v-icon icon="mdi-file-document-outline" color="primary" class="mr-2" />
        最优编译.md
      </v-card-title>
      
      <v-divider />
      
      <v-card-text>
        <v-alert
          type="info"
          variant="tonal"
          class="mb-4"
          icon="mdi-information"
        >
          <strong>描述：</strong> 使用最优参数编译LaTeX文档
        </v-alert>
        
        <v-card class="pa-4" color="surface-variant" elevation="0">
          <div class="d-flex align-center mb-2">
            <span class="text-overline text-secondary">命令内容</span>
            <v-spacer />
            <v-btn
              size="small"
              variant="outlined"
              color="primary"
              prepend-icon="mdi-content-copy"
            >
              复制命令
            </v-btn>
          </div>
          <pre style="font-family: 'JetBrains Mono', monospace; font-size: 13px; line-height: 1.6; color: rgb(var(--v-theme-on-surface)); white-space: pre-wrap;">{{ commandContent.split('\\n').slice(2).join('\\n') }}</pre>
        </v-card>
        
        <v-card class="mt-4 pa-4" color="surface" elevation="0">
          <v-icon icon="mdi-lightbulb-outline" color="accent" class="mr-2" />
          <span class="text-caption text-secondary">
            提示：按 <kbd>~</kbd> 键快速唤醒命令库，输入命令名称快速搜索
          </span>
        </v-card>
      </v-card-text>
    </v-card>
  </div>
</template>

<style scoped>
pre {
  margin: 0;
}

kbd {
  background: rgb(var(--v-theme-background));
  padding: 2px 6px;
  border-radius: 4px;
  border: 1px solid rgb(var(--v-theme-border));
  font-family: monospace;
}
</style>
