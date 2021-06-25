package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

var jsonstr = `[
  [
    {
      "content": "class=\"link-url\">http://192.168.0.63/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 178,
      "endPos": 219
    },
    {
      "content": "http://192.168.0.63/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 195,
      "endPos": 215
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 214,
      "endPos": 215
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://192.168.0.55/eams/loginExt.action</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 823,
      "endPos": 884
    },
    {
      "content": "http://192.168.0.55/eams/loginExt.action",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 840,
      "endPos": 880
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/eams/loginExt.action",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 859,
      "endPos": 880
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://60.190.32.202:8888/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 1468,
      "endPos": 1515
    },
    {
      "content": "http://60.190.32.202:8888/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 1485,
      "endPos": 1511
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": ":8888/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 1505,
      "endPos": 1511
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://webscan.360.cn/index/checkwebsite/url/jwc.zjpu.edu.cn</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 2139,
      "endPos": 2220
    },
    {
      "content": "http://webscan.360.cn/index/checkwebsite/url/jwc.zjpu.edu.cn",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 2156,
      "endPos": 2216
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/index/checkwebsite/url/jwc.zjpu.edu.cn",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 2177,
      "endPos": 2216
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://xm.zjkjt.gov.cn/redirect.action</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 2874,
      "endPos": 2933
    },
    {
      "content": "http://xm.zjkjt.gov.cn/redirect.action",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 2891,
      "endPos": 2929
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/redirect.action",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 2913,
      "endPos": 2929
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://www.zjnsf.gov.cn/login.aspx?ReturnUrl=/b/home.aspx</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 3559,
      "endPos": 3637
    },
    {
      "content": "http://www.zjnsf.gov.cn/login.aspx?ReturnUrl=/b/home.aspx",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 3576,
      "endPos": 3633
    },
    {
      "content": "www.",
      "isParticipating": true,
      "groupNum": 2,
      "groupName": 2,
      "startPos": 3583,
      "endPos": 3587
    },
    {
      "content": "/login.aspx?ReturnUrl=/b/home.aspx",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 3599,
      "endPos": 3633
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://www.zjkysz.cn/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 4223,
      "endPos": 4265
    },
    {
      "content": "http://www.zjkysz.cn/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 4240,
      "endPos": 4261
    },
    {
      "content": "www.",
      "isParticipating": true,
      "groupNum": 2,
      "groupName": 2,
      "startPos": 4247,
      "endPos": 4251
    },
    {
      "content": "/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 4260,
      "endPos": 4261
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://xmsb.nbsti.gov.cn/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 4855,
      "endPos": 4901
    },
    {
      "content": "http://xmsb.nbsti.gov.cn/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 4872,
      "endPos": 4897
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 4896,
      "endPos": 4897
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://jwc.zjpc.net.cn/dfiles/11353/r/cms/jwc/default/images/er_dh2_bg.jpg</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 5540,
      "endPos": 5636
    },
    {
      "content": "https://jwc.zjpc.net.cn/dfiles/11353/r/cms/jwc/default/images/er_dh2_bg.jpg",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 5557,
      "endPos": 5632
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/dfiles/11353/r/cms/jwc/default/images/er_dh2_bg.jpg",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 5580,
      "endPos": 5632
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://kyc.zjpc.net.cn/jybrwshkxyjybxmjtcl.rar</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 6186,
      "endPos": 6254
    },
    {
      "content": "https://kyc.zjpc.net.cn/jybrwshkxyjybxmjtcl.rar",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 6203,
      "endPos": 6250
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/jybrwshkxyjybxmjtcl.rar",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 6226,
      "endPos": 6250
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://kyc.zjpc.net.cn/dfiles/szsktjtcl.rar</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 6794,
      "endPos": 6859
    },
    {
      "content": "https://kyc.zjpc.net.cn/dfiles/szsktjtcl.rar",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 6811,
      "endPos": 6855
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/dfiles/szsktjtcl.rar",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 6834,
      "endPos": 6855
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://kyc.zjpc.net.cn/sjkghktjtcl.rar</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 7394,
      "endPos": 7454
    },
    {
      "content": "https://kyc.zjpc.net.cn/sjkghktjtcl.rar",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 7411,
      "endPos": 7450
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/sjkghktjtcl.rar",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 7434,
      "endPos": 7450
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://kyc.zjpc.net.cn/dfiles/1.rar</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 7986,
      "endPos": 8043
    },
    {
      "content": "https://kyc.zjpc.net.cn/dfiles/1.rar",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 8003,
      "endPos": 8039
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/dfiles/1.rar",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 8026,
      "endPos": 8039
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://jwc.zjpc.net.cn/__local/4/D9/8F/ADA03740FA576E9E5520A233CFE_F46FADA1_7A69.zip?e=.zip</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 8631,
      "endPos": 8744
    },
    {
      "content": "https://jwc.zjpc.net.cn/__local/4/D9/8F/ADA03740FA576E9E5520A233CFE_F46FADA1_7A69.zip?e=.zip",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 8648,
      "endPos": 8740
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/__local/4/D9/8F/ADA03740FA576E9E5520A233CFE_F46FADA1_7A69.zip?e=.zip",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 8671,
      "endPos": 8740
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://jwc.zjpc.net.cn/__local/E/F4/BD/0223BF27C1BF86BB5D990D7754E_1AC79588_1480A.rar?e=.rar</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 9353,
      "endPos": 9467
    },
    {
      "content": "https://jwc.zjpc.net.cn/__local/E/F4/BD/0223BF27C1BF86BB5D990D7754E_1AC79588_1480A.rar?e=.rar",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 9370,
      "endPos": 9463
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/__local/E/F4/BD/0223BF27C1BF86BB5D990D7754E_1AC79588_1480A.rar?e=.rar",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 9393,
      "endPos": 9463
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://jwc.zjpc.net.cn/system/resource/js/photoswipe/3.0.5.1/icons-2x.png</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 10052,
      "endPos": 10147
    },
    {
      "content": "https://jwc.zjpc.net.cn/system/resource/js/photoswipe/3.0.5.1/icons-2x.png",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 10069,
      "endPos": 10143
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/system/resource/js/photoswipe/3.0.5.1/icons-2x.png",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 10092,
      "endPos": 10143
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://cms.zjpc.net.cn/moodle/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 10745,
      "endPos": 10796
    },
    {
      "content": "http://cms.zjpc.net.cn/moodle/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 10762,
      "endPos": 10792
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/moodle/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 10784,
      "endPos": 10792
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://xm.zjkjt.gov.cn/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 11295,
      "endPos": 11339
    },
    {
      "content": "http://xm.zjkjt.gov.cn/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 11312,
      "endPos": 11335
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 11334,
      "endPos": 11335
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://kyc.zjpc.net.cn/__local/D/BD/D5/9A1C9CD4C550020A259A175A966_589785BB_BE83.rar?e=.rar</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 11996,
      "endPos": 12109
    },
    {
      "content": "https://kyc.zjpc.net.cn/__local/D/BD/D5/9A1C9CD4C550020A259A175A966_589785BB_BE83.rar?e=.rar",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 12013,
      "endPos": 12105
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/__local/D/BD/D5/9A1C9CD4C550020A259A175A966_589785BB_BE83.rar?e=.rar",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 12036,
      "endPos": 12105
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://kyc.zjpc.net.cn/__local/A/4E/55/1446097CF14F56133320F083468_D9A21D33_59C53.rar?e=.rar</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 12718,
      "endPos": 12832
    },
    {
      "content": "https://kyc.zjpc.net.cn/__local/A/4E/55/1446097CF14F56133320F083468_D9A21D33_59C53.rar?e=.rar",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 12735,
      "endPos": 12828
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/__local/A/4E/55/1446097CF14F56133320F083468_D9A21D33_59C53.rar?e=.rar",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 12758,
      "endPos": 12828
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://192.168.0.162:8080/system/editor/ueditor_news/lang/zh-cn/images/image.gif</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 13423,
      "endPos": 13524
    },
    {
      "content": "http://192.168.0.162:8080/system/editor/ueditor_news/lang/zh-cn/images/image.gif",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 13440,
      "endPos": 13520
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": ":8080/system/editor/ueditor_news/lang/zh-cn/images/image.gif",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 13460,
      "endPos": 13520
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://jwc.zjpc.net.cn/system/editor/ueditor_news/themes/default/images/spacer.gif</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 14162,
      "endPos": 14266
    },
    {
      "content": "https://jwc.zjpc.net.cn/system/editor/ueditor_news/themes/default/images/spacer.gif",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 14179,
      "endPos": 14262
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/system/editor/ueditor_news/themes/default/images/spacer.gif",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 14202,
      "endPos": 14262
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://kyc.zjpc.net.cn/system/resource/js/photoswipe/3.0.5.1/icons-2x.png</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 14852,
      "endPos": 14947
    },
    {
      "content": "https://kyc.zjpc.net.cn/system/resource/js/photoswipe/3.0.5.1/icons-2x.png",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 14869,
      "endPos": 14943
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/system/resource/js/photoswipe/3.0.5.1/icons-2x.png",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 14892,
      "endPos": 14943
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://zsjy.zjpc.net.cn/medicine/index.php</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 15557,
      "endPos": 15620
    },
    {
      "content": "http://zsjy.zjpc.net.cn/medicine/index.php",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 15574,
      "endPos": 15616
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/medicine/index.php",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 15597,
      "endPos": 15616
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://sysx.zjpc.net.cn/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 16120,
      "endPos": 16165
    },
    {
      "content": "http://sysx.zjpc.net.cn/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 16137,
      "endPos": 16161
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 16160,
      "endPos": 16161
    }
  ],
  [
    {
      "content": "class=\"link-url\">http://kyxm.zjpc.net.cn/</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 16665,
      "endPos": 16710
    },
    {
      "content": "http://kyxm.zjpc.net.cn/",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 16682,
      "endPos": 16706
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 16705,
      "endPos": 16706
    }
  ],
  [
    {
      "content": "class=\"link-url\">https://zs.zjpc.net.cn/images/icon2_11.png</a>",
      "isParticipating": true,
      "groupNum": 0,
      "groupName": null,
      "startPos": 17227,
      "endPos": 17290
    },
    {
      "content": "https://zs.zjpc.net.cn/images/icon2_11.png",
      "isParticipating": true,
      "groupNum": 1,
      "groupName": 1,
      "startPos": 17244,
      "endPos": 17286
    },
    {
      "content": "",
      "isParticipating": false,
      "groupNum": 2,
      "groupName": 2,
      "startPos": -1,
      "endPos": -1
    },
    {
      "content": "/images/icon2_11.png",
      "isParticipating": true,
      "groupNum": 3,
      "groupName": 3,
      "startPos": 17266,
      "endPos": 17286
    }
  ]
]`
func main() {
	js, err := simplejson.NewJson([]byte(jsonstr))
	if err != nil{
		panic(err.Error())
	}
	jsarray,_ := js.Array()
	fmt.Println(jsarray[0].)
}