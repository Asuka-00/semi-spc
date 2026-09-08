/**
 * SPC 前端共享常量、标签与样式辅助
 */
export const EQP_TYPE_VALUES = ['LITHO', 'ETCH', 'CVD', 'PVD', 'IMP', 'DIFF', 'CMP', 'METROLOGY', 'OTHER']
export const CHART_TYPE_VALUES = ['I_MR', 'XBAR_R', 'XBAR_S', 'P', 'NP', 'C', 'U', 'EWMA', 'CUSUM']
export const DATA_TYPE_VALUES = ['VARIABLE', 'ATTRIBUTE']
export const SAMPLE_LEVEL_VALUES = ['LOT', 'WAFER', 'SITE']
export const LOT_TYPE_VALUES = ['PROD', 'ENG', 'PILOT']
export const SEVERITY_VALUES = ['INFO', 'WARN', 'CRIT']
export const ALARM_TYPE_VALUES = ['OOC', 'OOS']
export const ALARM_STATUS_VALUES = ['OPEN', 'ACK', 'CLOSED']
export const TRIGGER_VALUES = ['OOC', 'OOS', 'BOTH']
export const LIMIT_SOURCE_VALUES = ['CALC', 'MANUAL']
export const LIMIT_METHOD_VALUES = ['CALC', 'MANUAL']

export const TIMEZONES = [
  { label: 'Asia/Shanghai', value: 'Asia/Shanghai' },
  { label: 'Asia/Singapore', value: 'Asia/Singapore' },
  { label: 'UTC', value: 'UTC' }
]

export const optionOf = (t, prefix, values) =>
  values.map((value) => ({ value, label: t(`${prefix}.${value}`) }))

export const labelOf = (options, value, fallback = value || '-') => {
  const hit = options.find((item) => item.value === value)
  return hit ? hit.label : fallback
}

export const statusTag = (status) => (Number(status) === 1 ? 'success' : 'info')
export const statusText = (status, t) =>
  Number(status) === 1 ? t('common.enabled') : t('common.disabled')

export const alarmTypeTag = (type) => (type === 'OOS' ? 'danger' : 'warning')
export const severityTag = (sev) => {
  if (sev === 'CRIT') return 'danger'
  if (sev === 'WARN') return 'warning'
  return 'info'
}
export const alarmStatusTag = (status) => {
  if (status === 'OPEN') return 'danger'
  if (status === 'ACK') return 'warning'
  return 'success'
}
export const alarmStatusText = (status, t) => {
  if (status === 'OPEN') return t('spc.alarm.open')
  if (status === 'ACK') return t('spc.alarm.acked')
  if (status === 'CLOSED') return t('spc.alarm.closed')
  return status || '-'
}

export const capabilityLevel = (value, t) => {
  if (value == null) return { text: '-', type: 'info' }
  if (value >= 1.67) return { text: t('spc.capability.excellent'), type: 'success' }
  if (value >= 1.33) return { text: t('spc.capability.good'), type: 'warning' }
  if (value >= 1.0) return { text: t('spc.capability.fair'), type: 'info' }
  return { text: t('spc.capability.poor'), type: 'danger' }
}

export const num = (value, digits = 3) => {
  if (value === null || value === undefined || value === '') return '-'
  const n = Number(value)
  return Number.isFinite(n) ? n.toFixed(digits) : '-'
}

export const requiredRule = (t, trigger = 'blur') => [
  { required: true, message: t('common.pleaseInput'), trigger }
]
export const selectRule = (t) => [{ required: true, message: t('common.pleaseSelect'), trigger: 'change' }]
