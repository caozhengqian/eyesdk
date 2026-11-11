declare const monitor: {
  init(options?: Record<string, unknown>): void;
  errorBoundary(error: unknown, info?: string): void;
  performance(): void;
  error(): void;
  behavior(): void;
  install(Vue: unknown, options?: Record<string, unknown>): void;
};

export default monitor;
export const errorBoundary: (error: unknown, info?: string) => void;
export const init: (options?: Record<string, unknown>) => void;
export const install: (Vue: unknown, options?: Record<string, unknown>) => void;
