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

           const chart =  Highcharts.stockChart('container_' + t, {
                rangeSelector: {
                    enabled: true,
                    buttons: [{
                        type: 'day',
                        count: 1,
                        text: '1d'
                    }, {
                        type: 'week',
                        count: 1,
                        text: '1w'
                    }, {
                        type: 'month',
                        count: 1,
                        text: '1m'
                    },{
                        type: 'month',
                         count: 6,
                        text: '6m'
                    },{
                      type: 'year',
                       count: 1,
                      text: '1y'
                  }, {type: 'year',
                       count: 2,
                      text: '2y'
                  },{type: 'year',
                       count: 3,
                      text: '3y'
                  },{
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
                                    this.series.chart.xAxis[0].drawCrosshair({plotX: this.plotX});
                                    this.series.chart.yAxis[0].drawCrosshair({plotY: this.plotY});
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
                    offset: 25,
                    zIndex: 2 ,
                    height: '70%',
                    crosshair: {
                        snap: false
                    },
                    accessibility: {
                        description: 'price'
                    }
                }, {
                    top: '70%',
                    height: '30%',
                    accessibility: {
                        description: 'volume'
                    }
                }],

                tooltip: {
                    shared: true,
                    split: false,
                    useHTML: true,
                    shadow: false,
                    positioner: function () {
                        return {x: 50, y: 10};
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
                }]
            });
        })();
    }

document.addEventListener('DOMContentLoaded', function() {
    var today = new Date().toISOString().slice(0, 10);
    document.getElementById('dateInput').value = today;

    var toggleChartCheckbox = document.getElementById('toggleChart');

    // 找到按鈕和日期輸入框
    var setDefaultDateButton = document.getElementById('setDefaultDateButton');
    var dateInput = document.getElementById('dateInput');

    // 在按鈕點擊時設置日期輸入框的值為今天的日期
    setDefaultDateButton.addEventListener('click', function() {
        var timestamp = new Date(dateInput.value).getTime();

        var is = 0
        if (toggleChartCheckbox.checked) {
            is = 1
        }

        (async () => {
            const dataDay = await fetch(
                    'http://127.0.0.1:8080/kline?t=day&end='+timestamp + '&is='+is
                ).then(response => response.json());

             const dataWeek = await fetch(
                    'http://127.0.0.1:8080/kline?t=week&end='+timestamp+ '&is='+is
                ).then(response => response.json());

            const dataMonth = await fetch(
                'http://127.0.0.1:8080/kline?t=month&end='+timestamp+ '&is='+is
            ).then(response => response.json());

            updateChart('container_day', dataDay);
            updateChart('container_week', dataWeek);
            updateChart('container_month', dataMonth);
        })();
    });
});

function updateChart(chartId, data) {
    var chart = Highcharts.charts.find(function(chart) {
        return chart.renderTo.id === chartId;
    });

    chart.series[0].setData(data);
    chart.hideLoading();
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