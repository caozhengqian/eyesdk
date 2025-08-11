import { lazyReportBatch } from '../report';
export default function observerLoad () {
    window.addEventListener('pageShow', function (event) {
        requestAnimationFrame(() =>{
            //load事件都注册完了，DOMContentLoaded事件可能没注册完
            ['load'].forEach((type) => {
                const reportData = {
                    type: 'performance',
                    subType: type,
                    pageUrl: window.location.href,
                    startTime: performance.now()- event.timeStamp
                }
                // 发送数据
                lazyReportBatch(reportData);
            });

        }, true);
    });
}
