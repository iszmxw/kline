package main

func main() {


	//var hqUrl = '//hq.sinajs.cn/rn=%rn%list=';
	////   浠锋牸淇℃伅
	//var _quoteTemp = doT.template($("#quoteTmp").html());
	//var _cursymbol = _symbol || 'fx_seurusd';
	//
	//function _getQuoteInfo() {
	//	var url = hqUrl.replace('%rn%', Autil.random()) + _cursymbol;
	//	Autil.getJsData(url, function () {
	//		if (window['hq_str_' + _cursymbol]) {
	//			var d = window['hq_str_' + _cursymbol].split(',');
	//			var _data = {};
	//			var _unit = 4;
	//			if((d[3]*1).toFixed(4) == '0.0000') {
	//				_unit = 6;
	//			}
	//			_data.name = d[9].replace('\u5373\u671F\u6C47\u7387', '');
	//			_data.price = (1 * d[8]).toFixed(_unit);
	//			_data.spj = (1 * d[3]).toFixed(_unit);
	//			_data.open = (1 * d[5]).toFixed(_unit);
	//			_data.higner = (1 * d[6]).toFixed(_unit);
	//			_data.lower = (1 * d[7]).toFixed(_unit);
	//			_data.updownCount = (1 * (d[8] - d[3])).toFixed(_unit);
	//			_data.fullcompany = d[13];
	//			_data.company = Autil.cutstr(d[13], 0, 40);
	//			_data.percent = Autil.getPercent(d[8], d[3], 4).p;
	//			_data.updown = Autil.getPercent(d[8], d[3], _unit).updown;
	//			_data.zf = Autil.getPercent2(d[6] - d[7], d[3], _unit);
	//			_data.bofu = 1 * (d[6] - d[7]).toFixed(_unit);
	//			_data.date = d[17] + ' ' + d[0];
	//			document.getElementById('quoteWrap').innerHTML = _quoteTemp(_data);
	//		}
	//	})
	//}
	//
	//
	//(function anonymous(it) {
	//	var out=' <div class="price_bar '+(it.updown)+'"> <div class="price"><h5>'+(it.price)+'</h5><label></label></div> <ul class="change"> <li class="change_li">'+(it.updownCount)+' ('+(it.percent)+')</li> <li class="fontSmall">'+(it.date)+'</li> </ul> </div> <div class="quote_detail_wrap"> <ul class="clearfix"> <li><label>今开</label>'+(it.open)+'</li> <li><label>昨收</label>'+(it.spj)+'</li> <li class="l2"><label>振幅</label>'+(it.zf)+'</li> <li class="l2 right"><label>波幅</label>'+(it.bofu)+'</li> <li><label>最低</label>'+(it.lower)+'</li> <li><label>最高</label>'+(it.higner)+'</li> <li class="l3 right"><label>报价</label>'+(it.company)+'</li> </ul> </div> ';return out;
	//})
	//
	//
	//var hq_str_fx_seurusd="" +
	//	"12:00:09," +
	//	"1.104500," +
	//	"1.104900," +
	//	"1.104700," +
	//	"0," +
	//	"1.104500," +
	//	"1.104500," +
	//	"1.104500," +
	//	"1.104500," +
	//	"欧元兑美元即期汇率," +
	//	"-0.020000," +
	//	"-0.000200," +
	//	"0," +
	//	"Hong Kong Forex Investment. Hong Kong," +
	//	"1.063500," +
	//	"1.063500," +
	//	"-," +
	//	"2022-04-02";
	//url := fmt.Sprintf("//hq.sinajs.cn/rn=%vlist=fx_seurusd", time.Now().Unix())
	//fmt.Println(url)
}
