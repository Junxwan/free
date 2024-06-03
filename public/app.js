function k(t) {
    // 取得當前日期
    const dataURL = 'http://127.0.0.1:8080/kline';
    (async () => {
        // Load the dataset
        const data = await fetch(
            'http://127.0.0.1:8080/kline?t=' + t
        ).then(response => response.json());

        // split the data set into ohlc and volume
        const ohlc = [],
            volume = [],
            dataLength = data.length;

        let previousCandleClose = 0;
        for (let i = 0; i < dataLength; i++) {
            ohlc.push([
                data[i][0], // the date
                data[i][1], // open
                data[i][2], // high
                data[i][3], // low
                data[i][4] // close
            ]);

            volume.push({
                x: data[i][0], // the date
                y: data[i][5], // the volume
                color: data[i][4] > previousCandleClose ? '#466742' : '#a23f43',
                labelColor: data[i][4] > previousCandleClose ? '#51a958' : '#ea3d3d'
            });
            previousCandleClose = data[i][4];
        }

        Highcharts.setOptions({
            chart: {
                backgroundColor: '#0a0a0a'
            },
            title: {
                style: {
                    color: '#cccccc'
                }
            },
            xAxis: {
                gridLineColor: '#181816',
                labels: {
                    style: {
                        color: '#9d9da2'
                    }
                }
            },
            yAxis: {
                gridLineColor: '#181816',
                labels: {
                    style: {
                        color: '#9d9da2'
                    }
                }
            },
            tooltip: {
                backgroundColor: 'rgba(0, 0, 0, 0.5)',
                style: {
                    color: '#cdcdc9'
                }
            },
            scrollbar: {
                barBackgroundColor: '#464646',
                barBorderRadius: 0,
                barBorderWidth: 0,
                buttonBorderWidth: 0,
                buttonArrowColor: '#cccccc',
                rifleColor: '#cccccc',
                trackBackgroundColor: '#121211',
                trackBorderRadius: 0,
                trackBorderWidth: 1,
                trackBorderColor: '#464646'
            },
            exporting: {
                enabled: false
            },
            lang: {
                shortMonths: ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月'],
                weekdays: ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六']
            }
        });

        const chart = Highcharts.stockChart('container_' + t, {
            rangeSelector: {
                enabled: true,
                buttons: [{
                    type: 'day',
                    count: 1,
                    text: '1d'
                }, {
                    type: 'day',
                    count: 2,
                    text: '2d'
                }, {
                    type: 'day',
                    count: 3,
                    text: '3d'
                }, {
                    type: 'day',
                    count: 5,
                    text: '5d'
                }, {
                    type: 'week',
                    count: 1,
                    text: '1w'
                }, {
                    type: 'week',
                    count: 2,
                    text: '2w'
                }, {
                    type: 'month',
                    count: 1,
                    text: '1m'
                }, {
                    type: 'month',
                    count: 3,
                    text: '3m'
                }, {
                    type: 'month',
                    count: 6,
                    text: '6m'
                }, {
                    type: 'year',
                    count: 1,
                    text: '1y'
                }, {
                    type: 'year',
                    count: 2,
                    text: '2y'
                }, {
                    type: 'year',
                    count: 3,
                    text: '3y'
                },{
                                      type: 'year',
                                      count: 4,
                                      text: '4y'
                                  }, {
                    type: 'year',
                    count: 6,
                    text: '6y'
                }, {
                    type: 'all',
                    text: 'All'
                }],

                inputDateFormat: '%Y-%m-%d',
                inputEditDateFormat: '%Y-%m-%d'
            },

            navigator: {
                enabled: true,
                adaptToUpdatedData: true
            },
            title: {
                text: '台指期-' + t
            },

            plotOptions: {
                series: {
                    point: {
                        events: {
                            mouseOver: function () {
                                this.series.chart.xAxis[0].drawCrosshair({ plotX: this.plotX });
                                this.series.chart.yAxis[0].drawCrosshair({ plotY: this.plotY });
                            }
                        }
                    },
                    cursor: 'crosshair',
                    dataGrouping: {
                        enabled: false
                    },

                },
                candlestick: {
                    color: 'white',
                    upColor: 'red',
                    upLineColor: 'red',
                    lineColor: 'white'
                }
            },

            scrollbar: {
                enabled: true
            },

            xAxis: {
                minRange: 1,
                min: Date.UTC(2015, 0, 1),
                max: Date.now(),
                gridLineWidth: 1,
                crosshair: {
                    snap: false
                }
            },

            yAxis: [{
                offset: 40,
                zIndex: 2,
                height: '100%',
                crosshair: {
                    snap: false
                },
                accessibility: {
                    description: 'price'
                }
            }],

            tooltip: {
                shared: true,
                split: false,
                useHTML: true,
                shadow: false,
                positioner: function () {
                    return { x: 80, y: 80 };
                }
            },

            series: [{
                type: 'candlestick',
                zIndex: 1,
                id: 'price',
                name: 'AAPL Stock Price',
                data: ohlc,
                tooltip: {
                    valueDecimals: 2,
                    pointFormat: '<b>O</b> <span style="color: {point.color}">' +
                        '{point.open} </span>' +
                        '<b>H</b> <span style="color: {point.color}">' +
                        '{point.high}</span><br/>' +
                        '<b>L</b> <span style="color: {point.color}">{point.low} ' +
                        '</span>' +
                        '<b>C</b> <span style="color: {point.color}">' +
                        '{point.close}</span><br/>'
                }
            }, {
                type: 'sma',
                id: '5Ma',
                name: '5Ma',
                linkedTo: 'price',
                zIndex: 1,
                lineWidth: 0.5,
                color: '#ff8c00',
                params: {
                    period: 5
                },
                marker: {
                    enabled: false,
                    states: {
                        hover: {
                            enabled: false,
                        }
                    }
                }
            }, {
                type: 'sma',
                id: '10Ma',
                name: '10Ma',
                linkedTo: 'price',
                zIndex: 1,
                lineWidth: 0.5,
                color: '#00ffff',
                params: {
                    period: 10
                },
                marker: {
                    enabled: false,
                    states: {
                        hover: {
                            enabled: false,
                        }
                    }
                }
            }, {
                type: 'sma',
                id: '20Ma',
                name: '20Ma',
                linkedTo: 'price',
                zIndex: 1,
                lineWidth: 0.5,
                color: '#0a932f',
                params: {
                    period: 20
                },
                marker: {
                    enabled: false,
                    states: {
                        hover: {
                            enabled: false,
                        }
                    }
                }
            }, {
                type: 'sma',
                id: '60Ma',
                name: '60Ma',
                linkedTo: 'price',
                zIndex: 1,
                lineWidth: 0.5,
                color: '#d4b40f',
                params: {
                    period: 60
                },
                marker: {
                    enabled: false,
                    states: {
                        hover: {
                            enabled: false,
                        }
                    }
                }
            }, {
                type: 'sma',
                id: '120Ma',
                name: '120Ma',
                linkedTo: 'price',
                zIndex: 1,
                lineWidth: 0.5,
                color: '#d40f33',
                params: {
                    period: 120
                },
                marker: {
                    enabled: false,
                    states: {
                        hover: {
                            enabled: false,
                        }
                    }
                }
            }, {
                type: 'sma',
                id: '240Ma',
                name: '240Ma',
                linkedTo: 'price',
                zIndex: 1,
                lineWidth: 0.5,
                color: '#720fd4',
                params: {
                    period: 240
                },
                marker: {
                    enabled: false,
                    states: {
                        hover: {
                            enabled: false,
                        }
                    }
                }
            },
            {
                type: 'bb',
                id: 'bollingerBands',
                linkedTo: 'price',
                yAxis: 0,
                lineWidth: 0.5,
                color: '#0a932f',
                bottomLine: {
                    styles: {
                        lineColor: "gray",
                    },
                },
                topLine: {
                    styles: {
                        lineColor: "gray",
                    },
                },
            }],
        });

        setIndex(chart, t);
    })();
}

document.addEventListener('DOMContentLoaded', function () {
    const dateInput = document.getElementById('dateInput');
    const title = document.getElementById('dateTitle');

    dateInput.value = getDateExcludingWeekends().toISOString().slice(0, 10);

    const dayNames = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];

    var rightBtn = document.getElementById('right');
    var toggleChartCheckbox = document.getElementById('toggleChart');

    // 找到按鈕和日期輸入框
    var setDefaultDateButton = document.getElementById('setDefaultDateButton');

    // 在按鈕點擊時設置日期輸入框的值為今天的日期
    setDefaultDateButton.addEventListener('click', function () {
        if (toggleChartCheckbox.checked) {
            title.textContent = dateInput.value + " 夜盤 " + dayNames[(new Date(dateInput.value)).getDay()];
        } else {
            title.textContent = dateInput.value + " 日盤 " + dayNames[(new Date(dateInput.value)).getDay()];

            if (isThirdWednesdayOfMonth(Date(dateInput.value))) {
                title.textContent += " 月結算"
            }
        }

        updateK();
    });

    const yesterdayBtn = document.getElementById('yesterdayBtn');
    const tomorrowBtn = document.getElementById('tomorrowBtn');

    // 设置昨天的日期（跳过周末）
    yesterdayBtn.addEventListener('click', () => {
        if (toggleChartCheckbox.checked) {
            let currentDate = dateInput.value ? new Date(dateInput.value) : new Date();
            currentDate = getValidDate(currentDate, -1);
            dateInput.value = formatDate(currentDate);

            toggleChartCheckbox.checked = false;

            title.textContent = dateInput.value + " 日盤 " + dayNames[(new Date(dateInput.value)).getDay()];

            if (isThirdWednesdayOfMonth(Date(dateInput.value))) {
                title.textContent += " 月結算"
            }
        } else {
            toggleChartCheckbox.checked = true;

            title.textContent = dateInput.value + " 夜盤 " + dayNames[(new Date(dateInput.value)).getDay()];
        }

        updateK();
    });

    // 设置明天的日期（跳过周末）
    tomorrowBtn.addEventListener('click', () => {
        if (toggleChartCheckbox.checked) {
            toggleChartCheckbox.checked = false;

            title.textContent = dateInput.value + " 日盤 " + dayNames[(new Date(dateInput.value)).getDay()];
        } else {
            let currentDate = dateInput.value ? new Date(dateInput.value) : new Date();
            currentDate = getValidDate(currentDate, 1);
            dateInput.value = formatDate(currentDate);

            toggleChartCheckbox.checked = true;

            title.textContent = formatDate(currentDate) + " 夜盤 " + dayNames[(new Date(dateInput.value)).getDay()];
        }

        updateK();
    });

    rightBtn.addEventListener('click', () => {
        const inx = document.getElementById('index');
        if (inx.value === "1") {
            moveRight("5", "container_5");
            moveRight("30", "container_30");
            moveRight("60", "container_60");
        } else {
            moveRight("day", "container_day");
            moveRight("week", "container_week");
            moveRight("month", "container_month");
        }
    });

    // 获取小视窗元素
    var windowElement = document.getElementById('floating-window');

    // 记录鼠标按下时鼠标相对小视窗的位置
    var offsetX, offsetY;

    // 当鼠标按下时触发
    windowElement.addEventListener('mousedown', function (event) {
        // 计算鼠标相对小视窗的位置
        offsetX = event.clientX - windowElement.getBoundingClientRect().left;
        offsetY = event.clientY - windowElement.getBoundingClientRect().top;

        // 添加鼠标移动事件监听器
        document.addEventListener('mousemove', onMouseMove);
    });

    // 当鼠标释放时触发
    document.addEventListener('mouseup', function () {
        // 移除鼠标移动事件监听器
        document.removeEventListener('mousemove', onMouseMove);
    });

    // 鼠标移动时触发的函数
    function onMouseMove(event) {
        // 计算小视窗新的位置
        var newX = event.clientX - offsetX;
        var newY = event.clientY - offsetY;

        // 设置小视窗的新位置
        windowElement.style.left = newX + 'px';
        windowElement.style.top = newY + 'px';
    }
});

function updateChart(t, chartId, data) {
    console.log("updateChart " + t)

    var chart = Highcharts.charts.find(function (chart) {
        return chart.renderTo.id === chartId;
    });

    chart.series[0].setData(data);

    chart.series[1].update({
        params: {
            period: 5
        }
    }, true);

    chart.series[2].update({
        params: {
            period: 10
        }
    }, true);

    chart.series[3].update({
        params: {
            period: 20
        }
    }, true);

    chart.series[4].update({
        params: {
            period: 60
        }
    }, true);

    chart.series[5].update({
        params: {
            period: 120
        }
    }, true);

    chart.series[6].update({
        params: {
            period: 240
        }
    }, true);

    chart.hideLoading();

    // 获取图表的数据序列
    var series = chart.series[0]; // 假设您的 K 线图数据是第一个序列

   console.log(series.data.length)
    console.log(series.data)

    // 获取最后一个数据点
    var lastPoint = series.data[series.data.length - 1];

    // 获取最后一个数据点的时间戳
    var latestTimestamp = lastPoint.x;

    chart.xAxis[0].setExtremes(null, latestTimestamp);

    setIndex(chart, t);
}

function moveRight(t, chartId) {
    var chart = Highcharts.charts.find(function (chart) {
        return chart.renderTo.id === chartId;
    });

    // 获取图表的数据序列
    var series = chart.series[0]; // 假设您的 K 线图数据是第一个序列

    // 获取最后一个数据点
    var lastPoint = series.data[series.data.length - 1];

    // 获取最后一个数据点的时间戳
    var latestTimestamp = lastPoint.x;

    chart.xAxis[0].setExtremes(null, latestTimestamp);

    setIndex(chart, t);
}

function setIndex(chart, t) {
    switch (t) {
        case "day":
            // 3m
            chart.rangeSelector.clickButton(7, true);
            break;
        case "week":
            // 6m
            chart.rangeSelector.clickButton(document.getElementById('container_week').getAttribute('data-index'), true);
            break;
        case "month":
            // 2y
            chart.rangeSelector.clickButton(document.getElementById('container_month').getAttribute('data-index'), true);
            break;
        case "30":
            // 5d
            chart.rangeSelector.clickButton(3, true);
            break;
        case "60":
            // 2w
            chart.rangeSelector.clickButton(5, true);
            break;
        case "5":
            // 3d
            chart.rangeSelector.clickButton(2, true);
            break;
    }
}

function afterSetExtremes(e) {
    const { chart } = e.target;
    chart.showLoading('Loading data from server...');
    fetch(`${dataURL}?start=${Math.round(e.min)}&end=${Math.round(e.max)}`)
        .then(res => res.ok && res.json())
        .then(data => {
            chart.series[0].setData(data);
            chart.hideLoading();
        }).catch(error => console.error(error.message));
}

// 辅助函数：检查是否为周末
function isWeekend(date) {
    const day = date.getDay();
    return day === 0 || day === 6;
}

// 获取有效日期（非周末）
function getValidDate(date, offset) {
    date.setDate(date.getDate() + offset);
    while (isWeekend(date)) {
        date.setDate(date.getDate() + offset);
    }
    return date;
}

function formatDate(date) {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
}

function updateK() {
    rect();

    const dateInput = document.getElementById('dateInput');
    const toggleChartCheckbox = document.getElementById('toggleChart');
    const inx = document.getElementById('index');
    var timestamp = new Date(dateInput.value).getTime();

    var is = 0
    if (toggleChartCheckbox.checked) {
        is = 1
    }

    if (inx.value === "1") {
        (async () => {
            const data60 = await fetch(
                'http://127.0.0.1:8080/kline?t=60&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            const data30 = await fetch(
                'http://127.0.0.1:8080/kline?t=30&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            const data5 = await fetch(
                'http://127.0.0.1:8080/kline?t=5&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            updateChart("60", 'container_60', data60);
            updateChart("30", 'container_30', data30);
            updateChart("5", 'container_5', data5);
        })();
    }

    if (inx.value === "0") {
        (async () => {
            const dataDay = await fetch(
                'http://127.0.0.1:8080/kline?t=day&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            const dataWeek = await fetch(
                'http://127.0.0.1:8080/kline?t=week&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            const dataMonth = await fetch(
                'http://127.0.0.1:8080/kline?t=month&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            updateChart("day", 'container_day', dataDay);
            updateChart("week", 'container_week', dataWeek);
            updateChart("month", 'container_month', dataMonth);
        })();
    }

    if (inx.value === "m") {
        (async () => {
            const dataMonth = await fetch(
                'http://127.0.0.1:8080/kline?t=month&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            updateChart("month", 'container_month', dataMonth);
        })();
    }

     if (inx.value === "w") {
        (async () => {
            const dataWeek = await fetch(
                'http://127.0.0.1:8080/kline?t=week&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            updateChart("week", 'container_week', dataWeek);
        })();
    }

    if (inx.value === "mw") {
        (async () => {
            const dataWeek = await fetch(
                'http://127.0.0.1:8080/kline?t=week&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            const dataMonth = await fetch(
                'http://127.0.0.1:8080/kline?t=month&end=' + timestamp + '&is=' + is
            ).then(response => response.json());

            updateChart("week", 'container_week', dataWeek);
            updateChart("month", 'container_month', dataMonth);
        })();
    }
}

function getDateExcludingWeekends() {
    let date = new Date();
    let dayNumber = date.getDay();

    // 如果輸入的日期是星期六（6），則加兩天，變成星期一
    if (dayNumber === 6) {
        date.setDate(date.getDate() - 1);
    }
    // 如果輸入的日期是星期日（0），則加一天，變成星期一
    else if (dayNumber === 0) {
        date.setDate(date.getDate() - 2);
    }

    return date;
}

function isThirdWednesdayOfMonth(date) {
    // 確保輸入的日期有效
    if (!(date instanceof Date && !isNaN(date))) {
        return false;
    }

    // 獲取日期的年份和月份
    let year = date.getFullYear();
    let month = date.getMonth();

    // 設置日期為該月的第一天
    let firstDayOfMonth = new Date(year, month, 1);

    // 獲取第一天是星期幾
    let firstDayOfWeek = firstDayOfMonth.getDay();

    // 獲取第一個周三的日期
    let firstWednesdayDate = new Date(firstDayOfMonth);
    firstWednesdayDate.setDate(firstWednesdayDate.getDate() + (3 - firstDayOfWeek + 7) % 7);

    // 獲取第三周的周三的日期
    let thirdWednesdayDate = new Date(firstWednesdayDate);
    thirdWednesdayDate.setDate(thirdWednesdayDate.getDate() + 14);

    // 判斷輸入的日期是否為第三周的周三
    return date.getDate() === thirdWednesdayDate.getDate();
}

function rect() {
    var annotationShapes = document.querySelectorAll('.highcharts-annotation-shapes');

    // 遍历每个具有 highcharts-annotation-shapes 类的元素
    annotationShapes.forEach(function (element) {
        // 在当前元素下查找第一个 path 元素，并更改其填充颜色
        var firstPath = element.querySelector('path:first-of-type');
        if (firstPath) {
            firstPath.setAttribute('fill', 'rgba(0, 0, 0, 0)');
            firstPath.setAttribute('stroke-width', '1');
        }

        // 在当前元素下查找第二个 path 元素，并更改其填充颜色
        var secondPath = element.querySelector('path:nth-of-type(2)');
        if (secondPath) {
            secondPath.setAttribute('stroke', 'rgba(128, 0, 128, 100)');
            secondPath.setAttribute('stroke-width', '1');
        }
    });
}