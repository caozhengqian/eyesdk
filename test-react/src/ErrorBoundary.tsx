import { Component, type ReactNode } from 'react';

interface Props {
  fallback: ReactNode;
  children?: ReactNode;
}

interface State {
  hasError: boolean;
}

class ErrorBoundary extends Component<Props, State> {
  state: State = { hasError: false };

  // 捕获子组件渲染阶段抛出的错误，更新状态以便下一次渲染显示后备 UI。
  // 错误上报已交由 createRoot 的 onCaughtError / onUncaughtError 处理。
  static getDerivedStateFromError(_error: unknown): State {
    return { hasError: true };
  }

  render() {
    if (this.state.hasError) {
      return this.props.fallback;
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
