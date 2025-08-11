# 前端sdk
## 大纲图片
![alt text](READEME-img/sdk大纲.png)
## 已有收费产品对比
### 比较有名的产品
1. sentry： https://sentry.io/welcome/
2. 神策：https://www.sensorsdata.cn/
3. 听云：https://www.tingyun.com/
### 监控平台包括三个部分
1. 数据采集与上报
2. 数据存储和分析
3. 数据展示、数据报警和监控
### 设计思路
![alt text](READEME-img/设计思路.jpg)
### 控制台统计代码
1. FCP ≤ 1秒:首次内容绘制时间 (First Contentful Paint)
2. LCP ≤2 秒:最大内容绘制时间 (Largest Contentful Paint)
3. FID (first input delay) 
2. 首次内容绘制时间,FCP,LCP
```js
 const entryHandler = (list) => {
        for (const entry of list.getEntries()) {
            if (entry.name === 'first-contentful-paint') {
                observer.disconnect();
                const json = entry.toJSON();
                console.log(json);
            }
        }
    
    }
    // 统计和计算fp的时间
    const observer = new PerformanceObserver(entryHandler);
    // buffered: true 确保观察到所有paint事件
    observer.observe({type: 'paint', buffered: true});
```