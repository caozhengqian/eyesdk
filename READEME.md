## js监控
### 代码解释
1. FP:src/performace/observerPaint
2. FCP:src/performace/observerFCP
3. LCP:src/performace/observeLCP
4. 静态资源加载：LCP:src/performace/observerEntries
### 一、性能监控
#### 1、FP(`src/performace/observeLCP`)
1.  new PerformanceObserver(entryHandler)
2.  FP: entryHandler方法里获取`first-paint`（`src/performace/observerPaint`）
#### 2、FCP(`src/performace/observeLCP`)
1.  new PerformanceObserver(entryHandler)
2.  FCP: entryHandler方法里获取`first-contentful-paint`（`src/performace/observerFCP`）
#### 3、LCP(`src/performace/observeLCP`)
1.  new PerformanceObserver(entryHandler)
2.  LCP: entryHandler方法里list.getEntries(`src/performace/observeLCP`)
#### 4、静态资源加载(`src/performace/observerEntries`)
1.  new PerformanceObserver(entryHandler)
2.  entryHandler方法里获取资源的各种属性(名称、大小、加载时间等)。
#### 5、ajax请求消耗
1. xhr监听`src/performace/xhr`
   1. 获取XMLHttpRequest.prototype，并在原生的.send方法中增加监听
   2. 监听接口加载停止后触发的loadend（MDN浏览器原生方法）事件中，获取xhr中的接口url\开始结束事件等
2. fetch监听`src/performace/fetch`
   1. 重写window.fetch
### 二、错误监控`src/error/index`
#### 1、资源加载错误
1. window.addEventListener('error',function () {})
#### 2、Javascript错误
1. window.onerror = function (msg, url, lineNo, columnNo, error) {}
2. 语法错误（编译时会发现）
3. 运行时错误（可用try catch抑制）
4. 逻辑错误（错误的判断条件，很难发现）
#### 3、async await、promise错误
1. window.addEventListener('unhandledrejection',function () {})
2. try catch无法捕获`异步`
3. window.onerror捕获`异步+同步`，无法捕获`图片、js、css`加载错误 + `fetch和new Image`错误
4. window.addEventListener("error")能捕获`图片、js、css`加载错误，不能捕获`fetch和new Image`错误
#### 4、react错误捕获
### 三、行为监控
#### 1、PV、UV
#### 2、页面停留时间
#### 3、用户行为（click等）
### 四、上报的方法
#### 1、xhr
#### 2、image、gif
#### 3、sendBeacon
### 五、上报时机
#### 1、requestidleCallback
#### 2、setTimeOut
#### 3、beforeUnload
#### 4、缓存批量