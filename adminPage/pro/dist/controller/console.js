/** layuiAdmin.pro-v1.4.0 LPPL License By https://www.layui.com/admin/ */
; layui.define(function (e) {
    layui.use(["admin", "carousel"],
        function () {
            var e = layui.$,
                t = (layui.admin, layui.carousel),
                a = layui.element, i = layui.device();
            e(".layadmin-carousel").each(function () {
                var a = e(this); t.render({
                    elem: this, width: "100%", arrow: "none", interval: a.data("interval"),
                    autoplay: a.data("autoplay") === !0, trigger: i.ios || i.android ? "click" : "hover", anim: a.data("anim")
                })
            }),
                a.render("progress")
        }),
        layui.use(["admin", "carousel", "echarts"], function () {
            var pv=[],uv=[],ip=[],times= [],counts=[]
            $.ajax({
                type: "GET",
                url: "http://127.0.0.1:8080/api/web/flowHour",
                headers:{token:layui.data('layuiAdmin').token},
                success:function(res){
                    if (res.code==200){
                        pv = res.data.hour.pv
                        uv = res.data.hour.uv
                        ip = res.data.hour.ip
                        times = res.data.week
                        counts = res.data.count
                        var e = layui.$, t = layui.admin, a = layui.carousel, i = layui.echarts, l = [],
                n = [{
                    title: { text: "今日流量趋势", x: "center", textStyle: { fontSize: 14 } },
                    tooltip: { trigger: "axis" }, legend: { data: ["", ""] },
                    xAxis: [{ type: "category", boundaryGap: !1, data: ["00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00", "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"] }],
                    yAxis: [{ type: "value" }],
                    series: [
                        {
                            name: "PV",
                            type: "line",
                            smooth: !0, itemStyle: { normal: { areaStyle: { type: "default" } } },
                            data: pv
                        },
                        {
                            name: "UV",
                            type: "line",
                            smooth: !0, itemStyle: { normal: { areaStyle: { type: "default" } } },
                            data: uv
                        },
                        {
                            name: "IP",
                            type: "line",
                            smooth: !0, itemStyle: { normal: { areaStyle: { type: "default" } } },
                            data: ip
                        }]
                },
                {
                    title: { text: "最近一周新用户量", x: "center", textStyle: { fontSize: 14 } },
                    tooltip: { trigger: "axis", formatter: "{b}<br>新用户：{c}" },
                    xAxis: [{ type: "category", data: times }],
                    yAxis: [{ type: "value" }],
                    series: [{ type: "line", data: counts }]
                }],
                r = e("#LAY-index-dataview").children("div"),
                o = function (e) {
                    l[e] = i.init(r[e],layui.echartsTheme),
                        l[e].setOption(n[e]),
                        t.resize(function () { l[e].resize() })
                };

            if (r[0]) {
                o(0); var d = 0;
                a.on("change(LAY-index-dataview)",
                    function (e) { o(d = e.index) }),
                    layui.admin.on("side", function () { setTimeout(function () { o(d) }, 300) }),
                    layui.admin.on("hash(tab)", function () { layui.router().path.join("") || o(d) })
            }
                    }
                }   
            })



            
        }),
        layui.use("table", function () { var e = (layui.$, layui.table); e.render({ elem: "#LAY-index-topSearch", url: "./json/console/top-search.js", page: !0, cols: [[{ type: "numbers", fixed: "left" }, { field: "keywords", title: "关键词", minWidth: 300, templet: '<div><a href="https://www.baidu.com/s?wd={{ d.keywords }}" target="_blank" class="layui-table-link">{{ d.keywords }}</div>' }, { field: "frequency", title: "搜索次数", minWidth: 120, sort: !0 }, { field: "userNums", title: "用户数", sort: !0 }]], skin: "line" }), e.render({ elem: "#LAY-index-topCard", url: "./json/console/top-card.js", page: !0, cellMinWidth: 120, cols: [[{ type: "numbers", fixed: "left" }, { field: "title", title: "标题", minWidth: 300, templet: '<div><a href="{{ d.href }}" target="_blank" class="layui-table-link">{{ d.title }}</div>' }, { field: "username", title: "发帖者" }, { field: "channel", title: "类别" }, { field: "crt", title: "点击率", sort: !0 }]], skin: "line" }) }), e("console", {})
});