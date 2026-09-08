import { createI18n } from 'vue-i18n'
import zhCN from './locales/zh-CN'
import enUS from './locales/en-US'
import jaJP from './locales/ja-JP'

export const SUPPORTED_LOCALES = [
  { value: 'zh-CN', label: '简体中文', short: '中' },
  { value: 'en-US', label: 'English', short: 'EN' },
  { value: 'ja-JP', label: '日本語', short: '日' }
]

export const DEFAULT_LOCALE = 'zh-CN'

export function readLocale() {
  const saved = localStorage.getItem('spc-locale')
  if (SUPPORTED_LOCALES.some((item) => item.value === saved)) return saved
  const nav = navigator.language || ''
  if (nav.startsWith('ja')) return 'ja-JP'
  if (nav.startsWith('en')) return 'en-US'
  return DEFAULT_LOCALE
}

export const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  locale: readLocale(),
  fallbackLocale: {
    zh: ['zh-CN'],
    en: ['en-US'],
    ja: ['ja-JP'],
    default: [DEFAULT_LOCALE]
  },
  messages: {
    'zh-CN': zhCN,
    zh: zhCN,
    'en-US': enUS,
    en: enUS,
    'ja-JP': jaJP,
    ja: jaJP
  }
})

export function setLocale(locale) {
  const next = SUPPORTED_LOCALES.some((item) => item.value === locale) ? locale : DEFAULT_LOCALE
  i18n.global.locale.value = next
  localStorage.setItem('spc-locale', next)
  document.documentElement.setAttribute('lang', next)
}

if (typeof document !== 'undefined') {
  document.documentElement.setAttribute('lang', readLocale())
}

export function translateMenuTitle(title, name) {
  const { t, te } = i18n.global
  if (name && te(`menu.${name}`)) return t(`menu.${name}`)
  if (title && te(`menuTitle.${title}`)) return t(`menuTitle.${title}`)
  return title
}

export function translateViolation(item) {
  if (!item) return ''
  const { t, te } = i18n.global
  const params = { ...(item.params || {}) }
  if (params.side && te(`spc.rule.side.${params.side}`)) {
    params.side = t(`spc.rule.side.${params.side}`)
  }
  if (params.trend && te(`spc.rule.trend.${params.trend}`)) {
    params.trend = t(`spc.rule.trend.${params.trend}`)
  }
  if (item.messageKey && te(item.messageKey)) {
    return t(item.messageKey, params)
  }
  return item.message || ''
}
