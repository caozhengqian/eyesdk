import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import ErrorBoundary from './ErrorBoundary'
import monitor from './monitor'

createRoot(document.getElementById('root')!, {
  // React 19 新的错误上报 API，替代 ErrorBoundary 中的 componentDidCatch 上报方式。
  // 组件错误被 ErrorBoundary 捕获后触发。
  onCaughtError(error, errorInfo) {
    monitor.errorBoundary(error, errorInfo.componentStack);
  },
  // 未被任何 ErrorBoundary 捕获、冒泡到根节点的错误。
  onUncaughtError(error, errorInfo) {
    monitor.errorBoundary(error, errorInfo.componentStack);
  },
}).render(
  <StrictMode>
    <ErrorBoundary fallback={<div>出错了，请稍后重试</div>}>
      <App />
    </ErrorBoundary>
  </StrictMode>,
)
