<template>
  <section class="spc-hero">
    <div class="relative flex flex-col gap-3 lg:flex-row lg:items-end lg:justify-between">
      <div>
        <p class="spc-kicker">{{ kicker }}</p>
        <h1 class="spc-title">{{ title }}</h1>
        <p v-if="subtitle" class="spc-subtitle">{{ subtitle }}</p>
      </div>
      <div v-if="$slots.actions" class="flex flex-wrap items-center gap-2">
        <slot name="actions" />
      </div>
    </div>
  </section>
</template>

<script setup>
defineProps({
  kicker: { type: String, default: 'STATISTICAL PROCESS CONTROL' },
  title: { type: String, required: true },
  subtitle: { type: String, default: '' }
})
</script>
