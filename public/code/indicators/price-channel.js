!/**
 * Highstock JS v11.4.3 (2024-05-22)
 *
 * Indicator series type for Highcharts Stock
 *
 * (c) 2010-2024 Daniel Studencki
 *
 * License: www.highcharts.com/license
 */function(t){"object"==typeof module&&module.exports?(t.default=t,module.exports=t):"function"==typeof define&&define.amd?define("highcharts/indicators/price-channel",["highcharts","highcharts/modules/stock"],function(e){return t(e),t.Highcharts=e,t}):t("undefined"!=typeof Highcharts?Highcharts:void 0)}(function(t){"use strict";var e=t?t._modules:{};function i(t,e,i,o){t.hasOwnProperty(e)||(t[e]=o.apply(null,i),"function"==typeof CustomEvent&&window.dispatchEvent(new CustomEvent("HighchartsModuleLoaded",{detail:{path:e,module:t[e]}})))}i(e,"Stock/Indicators/ArrayUtilities.js",[],function(){return{getArrayExtremes:function(t,e,i){return t.reduce((t,o)=>[Math.min(t[0],o[e]),Math.max(t[1],o[i])],[Number.MAX_VALUE,-Number.MAX_VALUE])}}}),i(e,"Stock/Indicators/MultipleLinesComposition.js",[e["Core/Series/SeriesRegistry.js"],e["Core/Utilities.js"]],function(t,e){var i;let{sma:{prototype:o}}=t.seriesTypes,{defined:r,error:s,merge:n}=e;return function(t){let e=["bottomLine"],i=["top","bottom"],a=["top"];function l(t){return"plot"+t.charAt(0).toUpperCase()+t.slice(1)}function p(t,e){let i=[];return(t.pointArrayMap||[]).forEach(t=>{t!==e&&i.push(l(t))}),i}function h(){let t=this,e=t.pointValKey,i=t.linesApiNames,a=t.areaLinesNames,h=t.points,c=t.options,u=t.graph,d={options:{gapSize:c.gapSize}},f=[],m=p(t,e),y=h.length,g;if(m.forEach((t,e)=>{for(f[e]=[];y--;)g=h[y],f[e].push({x:g.x,plotX:g.plotX,plotY:g[t],isNull:!r(g[t])});y=h.length}),t.userOptions.fillColor&&a.length){let e=f[m.indexOf(l(a[0]))],i=1===a.length?h:f[m.indexOf(l(a[1]))],r=t.color;t.points=i,t.nextPoints=e,t.color=t.userOptions.fillColor,t.options=n(h,d),t.graph=t.area,t.fillGraph=!0,o.drawGraph.call(t),t.area=t.graph,delete t.nextPoints,delete t.fillGraph,t.color=r}i.forEach((e,i)=>{f[i]?(t.points=f[i],c[e]?t.options=n(c[e].styles,d):s('Error: "There is no '+e+' in DOCS options declared. Check if linesApiNames are consistent with your DOCS line names."'),t.graph=t["graph"+e],o.drawGraph.call(t),t["graph"+e]=t.graph):s('Error: "'+e+" doesn't have equivalent in pointArrayMap. To many elements in linesApiNames relative to pointArrayMap.\"")}),t.points=h,t.options=c,t.graph=u,o.drawGraph.call(t)}function c(t){let e,i=[],r=[];if(t=t||this.points,this.fillGraph&&this.nextPoints){if((e=o.getGraphPath.call(this,this.nextPoints))&&e.length){e[0][0]="L",i=o.getGraphPath.call(this,t),r=e.slice(0,i.length);for(let t=r.length-1;t>=0;t--)i.push(r[t])}}else i=o.getGraphPath.apply(this,arguments);return i}function u(t){let e=[];return(this.pointArrayMap||[]).forEach(i=>{e.push(t[i])}),e}function d(){let t=this.pointArrayMap,e=[],i;e=p(this),o.translate.apply(this,arguments),this.points.forEach(o=>{t.forEach((t,r)=>{i=o[t],this.dataModify&&(i=this.dataModify.modifyValue(i)),null!==i&&(o[e[r]]=this.yAxis.toPixels(i,!0))})})}t.compose=function(t){let o=t.prototype;return o.linesApiNames=o.linesApiNames||e.slice(),o.pointArrayMap=o.pointArrayMap||i.slice(),o.pointValKey=o.pointValKey||"top",o.areaLinesNames=o.areaLinesNames||a.slice(),o.drawGraph=h,o.getGraphPath=c,o.toYData=u,o.translate=d,t}}(i||(i={})),i}),i(e,"Stock/Indicators/PC/PCIndicator.js",[e["Stock/Indicators/ArrayUtilities.js"],e["Stock/Indicators/MultipleLinesComposition.js"],e["Core/Color/Palettes.js"],e["Core/Series/SeriesRegistry.js"],e["Core/Utilities.js"]],function(t,e,i,o,r){let{sma:s}=o.seriesTypes,{merge:n,extend:a}=r;class l extends s{getValues(e,i){let o,r,s,n,a,l,p;let h=i.period,c=e.xData,u=e.yData,d=u?u.length:0,f=[],m=[],y=[];if(!(d<h)){for(p=h;p<=d;p++)n=c[p-1],a=u.slice(p-h,p),o=((r=(l=t.getArrayExtremes(a,2,1))[1])+(s=l[0]))/2,f.push([n,r,o,s]),m.push(n),y.push([r,o,s]);return{values:f,xData:m,yData:y}}}}return l.defaultOptions=n(s.defaultOptions,{params:{index:void 0,period:20},lineWidth:1,topLine:{styles:{lineColor:i.colors[2],lineWidth:1}},bottomLine:{styles:{lineColor:i.colors[8],lineWidth:1}},dataGrouping:{approximation:"averages"}}),a(l.prototype,{areaLinesNames:["top","bottom"],nameBase:"Price Channel",nameComponents:["period"],linesApiNames:["topLine","bottomLine"],pointArrayMap:["top","middle","bottom"],pointValKey:"middle"}),e.compose(l),o.registerSeriesType("pc",l),l}),i(e,"masters/indicators/price-channel.src.js",[e["Core/Globals.js"]],function(t){return t})});