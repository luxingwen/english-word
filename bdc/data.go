package main

var wordlist []Word = []Word{
	Word{Name: "excuse", Desc: "v.原谅"},
	Word{Name: "me", Desc: "pron.我(宾格)"},
	Word{Name: "yes", Desc: "ad.是的"},
	Word{Name: "is", Desc: "v.be动词现在时第三人称单数"},
	Word{Name: "this", Desc: "pron.这"},
	Word{Name: "your", Desc: "你的，你们的"},
	Word{Name: "handbag", Desc: "n.(女用)手提包"},
	Word{Name: "pardon", Desc: "int.原谅，请再说一遍"},
	Word{Name: "it", Desc: "pron.它"},
	Word{Name: "thank you", Desc: "感谢你(们)"},
	Word{Name: "very much", Desc: "非常地"},
	Word{Name: "pen", Desc: "n.钢笔"},
	Word{Name: "pencil", Desc: "n.铅笔"},
	Word{Name: "book", Desc: "n.书"},
	Word{Name: "watch", Desc: "n.手表"},
	Word{Name: "coat", Desc: "n.上衣，外衣"},
	Word{Name: "dress", Desc: "n.连衣裙"},
	Word{Name: "skirt", Desc: "n.裙子"},
	Word{Name: "shirt", Desc: "n.衬衣"},
	Word{Name: "car", Desc: "n.小汽车"},
	Word{Name: "house", Desc: "n.房子"},
	Word{Name: "umbrella", Desc: "n.伞"},
	Word{Name: "please", Desc: "int.请"},
	Word{Name: "here", Desc: "ad.这里"},
	Word{Name: "my", Desc: "我的"},
	Word{Name: "ticket", Desc: "n.票"},
	Word{Name: "number", Desc: "n.号码"},
	Word{Name: "five", Desc: "num.五"},
	Word{Name: "sorry", Desc: "a.对不起的"},
	Word{Name: "sir", Desc: "n.先生"},
	Word{Name: "cloakroom", Desc: "n.衣帽存放处"},
	Word{Name: "suit", Desc: "n.一套衣服"},
	Word{Name: "school", Desc: "n.学校"},
	Word{Name: "teacher", Desc: "n.老师"},
	Word{Name: "son", Desc: "n.儿子"},
	Word{Name: "daughter", Desc: "n.女儿"},
	Word{Name: "Mr.", Desc: "先生"},
	Word{Name: "good", Desc: "a.好"},
	Word{Name: "morning", Desc: "n.早晨"},
	Word{Name: "Miss", Desc: "小姐"},
	Word{Name: "new", Desc: "a.新的"},
	Word{Name: "student", Desc: "n.学生"},
	Word{Name: "French", Desc: "a.& n.法国人"},
	Word{Name: "German", Desc: "a.& n.德国人"},
	Word{Name: "nice", Desc: "a.美好的"},
	Word{Name: "meet", Desc: "v.遇见"},
	Word{Name: "Japanese", Desc: "a.& n.日本人"},
	Word{Name: "Korean", Desc: "a.& n.韩国人"},
	Word{Name: "Chinese", Desc: "a.& n.中国人"},
	Word{Name: "too", Desc: "ad.也"},
	Word{Name: "make", Desc: "n.(产品的)牌号"},
	Word{Name: "Swedish", Desc: "a.瑞典的"},
	Word{Name: "English", Desc: "a.英国的"},
	Word{Name: "American", Desc: "a.美国的"},
	Word{Name: "Italian", Desc: "a.意大利的"},
	Word{Name: "Volvo", Desc: "n.沃尔沃"},
	Word{Name: "Peugeot", Desc: "n.标致"},
	Word{Name: "Mercedes", Desc: "n.梅赛德斯"},
	Word{Name: "Toyota", Desc: "n.丰田"},
	Word{Name: "Daewoo", Desc: "n.大宇"},
	Word{Name: "Mini", Desc: "n.迷你"},
	Word{Name: "Ford", Desc: "n.福特"},
	Word{Name: "Fiat", Desc: "n.菲亚特"},
	Word{Name: "I", Desc: "pron.我"},
	Word{Name: "am", Desc: "v.be动词现在时第一人称单数"},
	Word{Name: "are", Desc: "v.be动词现在时复数"},
	Word{Name: "name", Desc: "n.名字"},
	Word{Name: "what", Desc: "a.& pron.什么"},
	Word{Name: "nationality", Desc: "n.国籍"},
	Word{Name: "job", Desc: "n.工作"},
	Word{Name: "keyboard", Desc: "n.电脑键盘"},
	Word{Name: "operator", Desc: "n.操作人员"},
	Word{Name: "engineer", Desc: "n.工程师"},
	Word{Name: "policeman", Desc: "n.警察"},
	Word{Name: "policewoman", Desc: "n.女警察"},
	Word{Name: "taxi driver", Desc: "出租汽车司机"},
	Word{Name: "air hostess", Desc: "空中小姐"},
	Word{Name: "postman", Desc: "n.邮递员"},
	Word{Name: "nurse", Desc: "n.护士"},
	Word{Name: "mechanic", Desc: "n.机械师"},
	Word{Name: "hairdresser", Desc: "n.理发师"},
	Word{Name: "housewife", Desc: "n.家庭妇女"},
	Word{Name: "milkman", Desc: "n.送牛奶的人"},
	Word{Name: "hello", Desc: "int.喂(表示问候)"},
	Word{Name: "hi", Desc: "int.喂，嗨"},
	Word{Name: "how", Desc: "ad.怎样"},
	Word{Name: "today", Desc: "ad.今天"},
	Word{Name: "well", Desc: "a.身体好"},
	Word{Name: "fine", Desc: "a.美好的"},
	Word{Name: "thanks", Desc: "int.谢谢"},
	Word{Name: "goodbye", Desc: "int.再见"},
	Word{Name: "see", Desc: "v.见"},
	Word{Name: "fat", Desc: "a.胖的"},
	Word{Name: "woman", Desc: "n.女人"},
	Word{Name: "thin", Desc: "a.瘦的"},
	Word{Name: "tall", Desc: "a.高的"},
	Word{Name: "short", Desc: "a.矮的"},
	Word{Name: "dirty", Desc: "a.脏的"},
	Word{Name: "clean", Desc: "a.干净的"},
	Word{Name: "hot", Desc: "a.热的"},
	Word{Name: "cold", Desc: "a.冷的"},
	Word{Name: "old", Desc: "a.老的"},
	Word{Name: "young", Desc: "a.年轻的"},
	Word{Name: "busy", Desc: "a.忙的"},
	Word{Name: "lazy", Desc: "a.懒的"},
	Word{Name: "whose", Desc: "pron.谁的"},
	Word{Name: "blue", Desc: "a.蓝色的"},
	Word{Name: "perhaps", Desc: "ad.大概"},
	Word{Name: "white", Desc: "a.白色的"},
	Word{Name: "catch", Desc: "v.抓住"},
	Word{Name: "father", Desc: "n.父亲"},
	Word{Name: "mother", Desc: "n.母亲"},
	Word{Name: "blouse", Desc: "n.女衬衫"},
	Word{Name: "sister", Desc: "n.姐，妹"},
	Word{Name: "tie", Desc: "n.领带"},
	Word{Name: "brother", Desc: "n.兄，弟"},
	Word{Name: "his", Desc: "他的"},
	Word{Name: "her", Desc: "她的"},
	Word{Name: "colour", Desc: "n.颜色"},
	Word{Name: "green", Desc: "a.绿色"},
	Word{Name: "come", Desc: "v.来"},
	Word{Name: "upstairs", Desc: "ad.楼上"},
	Word{Name: "smart", Desc: "a.时髦的，巧妙的"},
	Word{Name: "hat", Desc: "n.帽子"},
	Word{Name: "same", Desc: "a.相同的"},
	Word{Name: "lovely", Desc: "a.可爱的，秀丽的"},
	Word{Name: "case", Desc: "n.箱子"},
	Word{Name: "carpet", Desc: "n.地毯"},
	Word{Name: "dog", Desc: "n.狗"},
	Word{Name: "customs", Desc: "n.海关"},
	Word{Name: "officer", Desc: "n.官员"},
	Word{Name: "girl", Desc: "n.女孩，姑娘"},
	Word{Name: "Danish", Desc: "a.& n.丹麦人"},
	Word{Name: "friend", Desc: "n.朋友"},
	Word{Name: "Norwegian", Desc: "a.& n.挪威人"},
	Word{Name: "passport", Desc: "n.护照"},
	Word{Name: "brown", Desc: "a.棕色的"},
	Word{Name: "tourist", Desc: "n.旅游者"},
	Word{Name: "Russian", Desc: "a.& n.俄罗斯人"},
	Word{Name: "Dutch", Desc: "a.& n.荷兰人"},
	Word{Name: "these", Desc: "pron.这些(this的复数)"},
	Word{Name: "red", Desc: "a.红色的"},
	Word{Name: "grey", Desc: "a.灰色的"},
	Word{Name: "yellow", Desc: "a.黄色的"},
	Word{Name: "black", Desc: "a.黑色的"},
	Word{Name: "orange", Desc: "a.橘黄色的"},
	Word{Name: "employee", Desc: "n.雇员"},
	Word{Name: "hard-working", Desc: "a.勤奋的"},
	Word{Name: "sales reps", Desc: "推销员"},
	Word{Name: "man", Desc: "n.男人"},
	Word{Name: "office", Desc: "n.办公室"},
	Word{Name: "assistant", Desc: "n.助手"},
	Word{Name: "matter", Desc: "n.事情"},
	Word{Name: "children", Desc: "n.孩子们(child 的复数)"},
	Word{Name: "tired", Desc: "a.累，疲乏"},
	Word{Name: "boy", Desc: "n.男孩"},
	Word{Name: "thirsty", Desc: "a.渴"},
	Word{Name: "mum", Desc: "n.妈妈(儿语)"},
	Word{Name: "sit down", Desc: "坐下"},
	Word{Name: "right", Desc: "a.好，可以"},
	Word{Name: "ice cream", Desc: "冰淇淋"},
	Word{Name: "big", Desc: "a.大的"},
	Word{Name: "small", Desc: "a.小的"},
	Word{Name: "open", Desc: "a.开着的"},
	Word{Name: "shut", Desc: "a.关着的"},
	Word{Name: "light", Desc: "a.轻的"},
	Word{Name: "heavy", Desc: "a.重的"},
	Word{Name: "long", Desc: "a.长的"},
	Word{Name: "shoe", Desc: "n.鞋子"},
	Word{Name: "grandfather", Desc: "n.祖父，外祖父"},
	Word{Name: "grandmother", Desc: "n.祖母，外祖母"},
	Word{Name: "give", Desc: "v.给"},
	Word{Name: "one", Desc: "pron.一个"},
	Word{Name: "which", Desc: "哪一个"},
	Word{Name: "empty", Desc: "a.空的"},
	Word{Name: "full", Desc: "a.满的"},
	Word{Name: "large", Desc: "a.大的"},
	Word{Name: "little", Desc: "a.小的"},
	Word{Name: "sharp", Desc: "a.尖的，锋利的"},
	Word{Name: "small", Desc: "a.小的"},
	Word{Name: "big", Desc: "a.大的"},
	Word{Name: "blunt", Desc: "a.钝的"},
	Word{Name: "box", Desc: "n.盒子，箱子"},
	Word{Name: "glass", Desc: "n.杯子"},
	Word{Name: "cup", Desc: "n.茶杯"},
	Word{Name: "bottle", Desc: "n.瓶子"},
	Word{Name: "tin", Desc: "n.罐头"},
	Word{Name: "knife", Desc: "n.刀子"},
	Word{Name: "fork", Desc: "n.叉子"},
	Word{Name: "on", Desc: "prep.在…之上"},
	Word{Name: "shelf", Desc: "n.架子，搁板"},
	Word{Name: "desk", Desc: "n.课桌"},
	Word{Name: "table", Desc: "n.桌子"},
	Word{Name: "plate", Desc: "n.盘子"},
	Word{Name: "cupboard", Desc: "n.食橱"},
	Word{Name: "cigarette", Desc: "n.香烟"},
	Word{Name: "television", Desc: "n.电视机"},
	Word{Name: "floor", Desc: "n.地板"},
	Word{Name: "dressing table", Desc: "梳妆台"},
	Word{Name: "magazine", Desc: "n.杂志"},
	Word{Name: "bed", Desc: "n.床"},
	Word{Name: "newspaper", Desc: "n.报纸"},
	Word{Name: "stereo", Desc: "n.立体声音响"},
	Word{Name: "Mrs.", Desc: "夫人"},
	Word{Name: "kitchen", Desc: "n.厨房"},
	Word{Name: "refrigerator", Desc: "n.电冰箱"},
	Word{Name: "right", Desc: "n.右边"},
	Word{Name: "electric", Desc: "a.带电的，可通电的"},
	Word{Name: "left", Desc: "n.左边"},
	Word{Name: "cooker", Desc: "n.炉子，炊具"},
	Word{Name: "middle", Desc: "n.中间"},
	Word{Name: "of", Desc: "prep.(属于)…的"},
	Word{Name: "room", Desc: "n.房间"},
	Word{Name: "cup", Desc: "n.杯子"},
	Word{Name: "where", Desc: "ad.在哪里"},
	Word{Name: "in", Desc: "prep.在…里"},
	Word{Name: "living room", Desc: "客厅"},
	Word{Name: "near", Desc: "prep.靠近"},
	Word{Name: "window", Desc: "n.窗户"},
	Word{Name: "armchair", Desc: "n.手扶椅"},
	Word{Name: "door", Desc: "n.门"},
	Word{Name: "picture", Desc: "n.图画"},
	Word{Name: "wall", Desc: "n.墙"},
	Word{Name: "trousers", Desc: "n.〔复数〕长裤"},
	Word{Name: "shut", Desc: "v.关门"},
	Word{Name: "bedroom", Desc: "n.卧室"},
	Word{Name: "untidy", Desc: "a.乱，不整齐"},
	Word{Name: "must", Desc: "modal verb.必须，应该"},
	Word{Name: "open", Desc: "v.打开"},
	Word{Name: "air", Desc: "v.使…通风，换换空气"},
	Word{Name: "put", Desc: "v.放置"},
	Word{Name: "clothes", Desc: "n.衣服"},
	Word{Name: "wardrobe", Desc: "n.大衣柜"},
	Word{Name: "dust", Desc: "v.掸掉灰尘"},
	Word{Name: "sweep", Desc: "v.扫"},
	Word{Name: "empty", Desc: "v.倒空，使…变空"},
	Word{Name: "read", Desc: "v.读"},
	Word{Name: "sharpen", Desc: "v.削尖，使锋利"},
	Word{Name: "put on", Desc: "穿上"},
	Word{Name: "take off", Desc: "脱掉"},
	Word{Name: "turn on", Desc: "开(电灯)"},
	Word{Name: "turn off", Desc: "关(电灯)"},
	Word{Name: "garden", Desc: "n.花园"},
	Word{Name: "under", Desc: "prep.在…之下"},
	Word{Name: "tree", Desc: "n.树"},
	Word{Name: "climb", Desc: "v.爬，攀登"},
	Word{Name: "who", Desc: "pron.谁"},
	Word{Name: "run", Desc: "v.跑"},
	Word{Name: "grass", Desc: "n.草，草地"},
	Word{Name: "after", Desc: "prep.在…之后"},
	Word{Name: "cat", Desc: "n.猫"},
	Word{Name: "type", Desc: "v.打字"},
	Word{Name: "letter", Desc: "n.信"},
	Word{Name: "basket", Desc: "n.篮子"},
	Word{Name: "eat", Desc: "v.吃"},
	Word{Name: "bone", Desc: "n.骨头"},
	Word{Name: "clean", Desc: "v.清洗"},
	Word{Name: "tooth", Desc: "n.牙齿"},
	Word{Name: "cook", Desc: "v.做(饭菜)"},
	Word{Name: "milk", Desc: "n.牛奶"},
	Word{Name: "meal", Desc: "n.饭，一顿饭"},
	Word{Name: "drink", Desc: "v.喝"},
	Word{Name: "tap", Desc: "n.(水)龙头"},
	Word{Name: "day", Desc: "n.日子"},
	Word{Name: "cloud", Desc: "n.云"},
	Word{Name: "sky", Desc: "n.天空"},
	Word{Name: "sun", Desc: "n.太阳"},
	Word{Name: "shine", Desc: "v.照耀"},
	Word{Name: "with", Desc: "prep.和…在一起"},
	Word{Name: "family", Desc: "n.家庭(成员)"},
	Word{Name: "walk", Desc: "v.走路，步行"},
	Word{Name: "over", Desc: "prep.跨越，在…之上"},
	Word{Name: "bridge", Desc: "n.桥"},
	Word{Name: "boat", Desc: "n.船"},
	Word{Name: "river", Desc: "n.河"},
	Word{Name: "ship", Desc: "n.轮船"},
	Word{Name: "aeroplane", Desc: "n.飞机"},
	Word{Name: "fly", Desc: "v.飞"},
	Word{Name: "sleep", Desc: "v.睡觉"},
	Word{Name: "shave", Desc: "v.刮脸"},
	Word{Name: "cry", Desc: "v.哭，喊"},
	Word{Name: "wash", Desc: "v.洗"},
	Word{Name: "wait", Desc: "v.等"},
	Word{Name: "jump", Desc: "v.跳"},
	Word{Name: "photograph", Desc: "n.照片"},
	Word{Name: "village", Desc: "n.村庄"},
	Word{Name: "valley", Desc: "n.山谷"},
	Word{Name: "between", Desc: "prep.在…之间"},
	Word{Name: "hill", Desc: "n.小山"},
	Word{Name: "another", Desc: "prep.另一个"},
	Word{Name: "wife", Desc: "n.妻子"},
	Word{Name: "along", Desc: "prep.沿着"},
	Word{Name: "bank", Desc: "n.河岸"},
	Word{Name: "water", Desc: "n.水"},
	Word{Name: "swim", Desc: "v.游泳"},
	Word{Name: "across", Desc: "prep.横过"},
	Word{Name: "building", Desc: "n.大楼，建筑物"},
	Word{Name: "park", Desc: "n.公园"},
	Word{Name: "into", Desc: "prep.进入"},
	Word{Name: "beside", Desc: "prep.在…旁"},
	Word{Name: "off", Desc: "prep.离开"},
	Word{Name: "work", Desc: "v.工作"},
	Word{Name: "hard", Desc: "ad.努力地"},
	Word{Name: "make", Desc: "v.做"},
	Word{Name: "bookcase", Desc: "n.书橱，书架"},
	Word{Name: "hammer", Desc: "n.锤子"},
	Word{Name: "paint", Desc: "v.上漆，涂"},
	Word{Name: "pink", Desc: "n.& a.粉红色"},
	Word{Name: "favourite", Desc: "a.最喜欢的"},
	Word{Name: "homework", Desc: "n.作业"},
	Word{Name: "listen", Desc: "v.听"},
	Word{Name: "dish", Desc: "n.盘子，碟子"},
	Word{Name: "front", Desc: "n.前面"},
	Word{Name: "in front of", Desc: "在…之前"},
	Word{Name: "careful", Desc: "a.小心的，仔细的"},
	Word{Name: "vase", Desc: "n.花瓶"},
	Word{Name: "drop", Desc: "v.掉下"},
	Word{Name: "flower", Desc: "n.花"},
	Word{Name: "show", Desc: "v.给…看"},
	Word{Name: "send", Desc: "v.送给"},
	Word{Name: "take", Desc: "v.带给"},
	Word{Name: "cheese", Desc: "n.乳酪，干酪"},
	Word{Name: "bread", Desc: "n.面包"},
	Word{Name: "soap", Desc: "n.肥皂"},
	Word{Name: "chocolate", Desc: "n.巧克力"},
	Word{Name: "sugar", Desc: "n.糖"},
	Word{Name: "coffee", Desc: "n.咖啡"},
	Word{Name: "tea", Desc: "n.茶"},
	Word{Name: "tobacco", Desc: "n.烟草，烟丝"},
	Word{Name: "bird", Desc: "n.鸟"},
	Word{Name: "any", Desc: "det.一些"},
	Word{Name: "some", Desc: "det.一些"},
	Word{Name: "of course", Desc: "当然"},
	Word{Name: "kettle", Desc: "n.水壶"},
	Word{Name: "behind", Desc: "prep.在…后面"},
	Word{Name: "teapot", Desc: "n.茶壶"},
	Word{Name: "now", Desc: "ad.现在，此刻"},
	Word{Name: "find", Desc: "v.找到"},
	Word{Name: "boil", Desc: "v.沸腾，开"},
	Word{Name: "can", Desc: "能够"},
	Word{Name: "boss", Desc: "n.老板，上司"},
	Word{Name: "minute", Desc: "n.分(钟)"},
	Word{Name: "ask", Desc: "v.请求，要求"},
	Word{Name: "handwriting", Desc: "n.书写"},
	Word{Name: "terrible", Desc: "a.糟糕的，可怕的"},
	Word{Name: "lift", Desc: "v.拿起，搬起，举起"},
	Word{Name: "cake", Desc: "n.饼，蛋糕"},
	Word{Name: "biscuit", Desc: "n.饼干"},
	Word{Name: "like", Desc: "v.喜欢，想要"},
	Word{Name: "want", Desc: "v.想"},
	Word{Name: "fresh", Desc: "a.新鲜的"},
	Word{Name: "egg", Desc: "n.鸡蛋"},
	Word{Name: "butter", Desc: "n.黄油"},
	Word{Name: "pure", Desc: "a.纯净的"},
	Word{Name: "honey", Desc: "n.蜂蜜"},
	Word{Name: "ripe", Desc: "a.成熟的"},
	Word{Name: "banana", Desc: "n.香蕉"},
	Word{Name: "jam", Desc: "n.果酱"},
	Word{Name: "sweet", Desc: "a.甜的"},
	Word{Name: "orange", Desc: "n.橙"},
	Word{Name: "Scotch whisky", Desc: "苏格兰威士忌"},
	Word{Name: "choice", Desc: "a.上等的，精选的"},
	Word{Name: "apple", Desc: "n.苹果"},
	Word{Name: "wine", Desc: "n.酒，果酒"},
	Word{Name: "beer", Desc: "n.啤酒"},
	Word{Name: "blackboard", Desc: "n.黑板"},
	Word{Name: "butcher", Desc: "n.卖肉的"},
	Word{Name: "meat", Desc: "n.(食用)肉"},
	Word{Name: "beef", Desc: "n.牛肉"},
	Word{Name: "lamb", Desc: "n.羔羊肉"},
	Word{Name: "husband", Desc: "n.丈夫"},
	Word{Name: "steak", Desc: "n.牛排"},
	Word{Name: "mince", Desc: "n.肉馅，绞肉"},
	Word{Name: "chicken", Desc: "n.鸡"},
	Word{Name: "tell", Desc: "v.告诉"},
	Word{Name: "truth", Desc: "n.实情"},
	Word{Name: "either", Desc: "ad.也(用于否定句)"},
	Word{Name: "tomato", Desc: "n.西红柿"},
	Word{Name: "potato", Desc: "n.土豆"},
	Word{Name: "cabbage", Desc: "n.卷心菜"},
	Word{Name: "lettuce", Desc: "n.莴苣"},
	Word{Name: "pea", Desc: "n.豌豆"},
	Word{Name: "bean", Desc: "n.豆角"},
	Word{Name: "pear", Desc: "n.梨"},
	Word{Name: "grape", Desc: "n.葡萄"},
	Word{Name: "peach", Desc: "n.桃"},
	Word{Name: "Greece", Desc: "n.希腊"},
	Word{Name: "climate", Desc: "n.气候"},
	Word{Name: "country", Desc: "n.国家"},
	Word{Name: "pleasant", Desc: "a.宜人的"},
	Word{Name: "weather", Desc: "n.天气"},
	Word{Name: "spring", Desc: "n.春季"},
	Word{Name: "windy", Desc: "a.有风的"},
	Word{Name: "warm", Desc: "a.温暖的"},
	Word{Name: "rain", Desc: "v.下雨"},
	Word{Name: "sometimes", Desc: "ad.有时"},
	Word{Name: "summer", Desc: "n.夏天"},
	Word{Name: "autumn", Desc: "n.秋天"},
	Word{Name: "winter", Desc: "n.冬天"},
	Word{Name: "snow", Desc: "v.下雪"},
	Word{Name: "January", Desc: "n.1月"},
	Word{Name: "February", Desc: "n.2月"},
	Word{Name: "March", Desc: "n.3月"},
	Word{Name: "April", Desc: "n.4月"},
	Word{Name: "May", Desc: "n.5月"},
	Word{Name: "June", Desc: "n.6月"},
	Word{Name: "July", Desc: "n.7月"},
	Word{Name: "August", Desc: "n.8月"},
	Word{Name: "September", Desc: "n.9月"},
	Word{Name: "October", Desc: "n.10月"},
	Word{Name: "November", Desc: "n.11月"},
	Word{Name: "December", Desc: "n.12月"},
	Word{Name: "the U.S.", Desc: "美国"},
	Word{Name: "Brazil", Desc: "n.巴西"},
	Word{Name: "Holland", Desc: "n.荷兰"},
	Word{Name: "England", Desc: "n.英国"},
	Word{Name: "France", Desc: "n.法国"},
	Word{Name: "Germany", Desc: "n.德国"},
	Word{Name: "Italy", Desc: "n.意大利"},
	Word{Name: "Norway", Desc: "n.挪威"},
	Word{Name: "Russia", Desc: "n.俄罗斯"},
	Word{Name: "Spain", Desc: "n.西班牙"},
	Word{Name: "Sweden", Desc: "n.瑞典"},
	Word{Name: "mild", Desc: "a.温和的，温暖的"},
	Word{Name: "always", Desc: "ad.总是"},
	Word{Name: "north", Desc: "n.北方"},
	Word{Name: "east", Desc: "n.东方"},
	Word{Name: "wet", Desc: "a.潮湿的"},
	Word{Name: "west", Desc: "n.西方"},
	Word{Name: "south", Desc: "n.南方"},
	Word{Name: "season", Desc: "n.季节"},
	Word{Name: "best", Desc: "ad.最"},
	Word{Name: "night", Desc: "n.夜晚"},
	Word{Name: "rise", Desc: "v.升起"},
	Word{Name: "early", Desc: "ad.早"},
	Word{Name: "set", Desc: "v.(太阳)落下去"},
	Word{Name: "late", Desc: "ad.晚，迟"},
	Word{Name: "interesting", Desc: "a.有趣的，有意思的"},
	Word{Name: "subject", Desc: "n.话题"},
	Word{Name: "conversation", Desc: "n.谈话"},
	Word{Name: "Australia", Desc: "n.澳大利亚"},
	Word{Name: "Australian", Desc: "n.澳大利亚人"},
	Word{Name: "Austria", Desc: "n.奥地利"},
	Word{Name: "Austrian", Desc: "n.奥地利人"},
	Word{Name: "Canada", Desc: "n.加拿大"},
	Word{Name: "Canadian", Desc: "n.加拿大人"},
	Word{Name: "China", Desc: "n.中国"},
	Word{Name: "Finland", Desc: "n.芬兰"},
	Word{Name: "Finnish", Desc: "n.芬兰人"},
	Word{Name: "India", Desc: "n.印度"},
	Word{Name: "Indian", Desc: "n.印度人"},
	Word{Name: "Japan", Desc: "n.日本"},
	Word{Name: "Nigeria", Desc: "n.尼日利亚"},
	Word{Name: "Nigerian", Desc: "n.尼日利亚人"},
	Word{Name: "Turkey", Desc: "n.土耳其"},
	Word{Name: "Turkish", Desc: "n.土耳其人"},
	Word{Name: "Korea", Desc: "n.韩国"},
	Word{Name: "Polish", Desc: "n.波兰人"},
	Word{Name: "Poland", Desc: "n.波兰"},
	Word{Name: "Thai", Desc: "n.泰国人"},
	Word{Name: "Thailand", Desc: "n.泰国"},
	Word{Name: "live", Desc: "v.住，生活"},
	Word{Name: "stay", Desc: "v.呆在，停留"},
	Word{Name: "home", Desc: "n.家 ad.在家，到家"},
	Word{Name: "housework", Desc: "n.家务"},
	Word{Name: "lunch", Desc: "n.午饭"},
	Word{Name: "afternoon", Desc: "n.下午"},
	Word{Name: "usually", Desc: "ad.通常"},
	Word{Name: "together", Desc: "ad.一起"},
	Word{Name: "evening", Desc: "n.晚上"},
	Word{Name: "arrive", Desc: "v.到达"},
	Word{Name: "night", Desc: "n.夜间"},
	Word{Name: "o'clock", Desc: "ad.点钟"},
	Word{Name: "shop", Desc: "n.商店"},
	Word{Name: "moment", Desc: "n.片刻，瞬间"},
	Word{Name: "envelope", Desc: "n.信封"},
	Word{Name: "writing paper", Desc: "信纸"},
	Word{Name: "shop assistant", Desc: "售货员"},
	Word{Name: "size", Desc: "n.尺寸，尺码，大小"},
	Word{Name: "pad", Desc: "n.信笺簿"},
	Word{Name: "glue", Desc: "n.胶水"},
	Word{Name: "chalk", Desc: "n.粉笔"},
	Word{Name: "change", Desc: "n.零钱，找给的钱"},
	Word{Name: "feel", Desc: "v.感觉"},
	Word{Name: "look", Desc: "v.看(起来)"},
	Word{Name: "must", Desc: "必须"},
	Word{Name: "call", Desc: "v.叫，请"},
	Word{Name: "doctor", Desc: "n.医生"},
	Word{Name: "telephone", Desc: "n.电话"},
	Word{Name: "remember", Desc: "v.记得，记住"},
	Word{Name: "mouth", Desc: "n.嘴"},
	Word{Name: "tongue", Desc: "n.舌头"},
	Word{Name: "bad", Desc: "a.坏的，严重的"},
	Word{Name: "cold", Desc: "n.感冒"},
	Word{Name: "news", Desc: "n.消息"},
	Word{Name: "headache", Desc: "n.头痛"},
	Word{Name: "aspirin", Desc: "n.阿斯匹林"},
	Word{Name: "earache", Desc: "n.耳痛"},
	Word{Name: "toothache", Desc: "n.牙痛"},
	Word{Name: "dentist", Desc: "n.牙医"},
	Word{Name: "stomach ache", Desc: "胃痛"},
	Word{Name: "medicine", Desc: "n.药"},
	Word{Name: "temperature", Desc: "n.温度"},
	Word{Name: "flu", Desc: "n.流行性感冒"},
	Word{Name: "measles", Desc: "n.麻疹"},
	Word{Name: "mumps", Desc: "n.腮腺炎"},
	Word{Name: "better", Desc: "a.形容词well的比较级"},
	Word{Name: "certainly", Desc: "ad.当然"},
	Word{Name: "get up", Desc: "起床"},
	Word{Name: "yet", Desc: "ad.还，仍"},
	Word{Name: "rich", Desc: "a.油腻的"},
	Word{Name: "food", Desc: "n.食物"},
	Word{Name: "remain", Desc: "v.保持，继续"},
	Word{Name: "play", Desc: "v.玩"},
	Word{Name: "match", Desc: "n.火柴"},
	Word{Name: "talk", Desc: "v.谈话"},
	Word{Name: "library", Desc: "n.图书馆"},
	Word{Name: "drive", Desc: "v.开车"},
	Word{Name: "so", Desc: "ad.如此地"},
	Word{Name: "quickly", Desc: "ad.快地"},
	Word{Name: "lean out of", Desc: "身体探出"},
	Word{Name: "break", Desc: "v.打破"},
	Word{Name: "Dad", Desc: "n.爸(儿语)"},
	Word{Name: "key", Desc: "n.钥匙"},
	Word{Name: "baby", Desc: "n.婴儿"},
	Word{Name: "hear", Desc: "v.听见"},
	Word{Name: "enjoy", Desc: "v.玩得快活"},
	Word{Name: "yourself", Desc: "pron.你自己"},
	Word{Name: "ourselves", Desc: "pron.我们自己"},
	Word{Name: "myself", Desc: "pron.我自己"},
	Word{Name: "themselves", Desc: "pron.他们自己"},
	Word{Name: "himself", Desc: "pron.他自己"},
	Word{Name: "herself", Desc: "pron.她自己"},
	Word{Name: "greengrocer", Desc: "n.蔬菜水果零售商"},
	Word{Name: "absent", Desc: "a.缺席的"},
	Word{Name: "Monday", Desc: "n.星期一"},
	Word{Name: "Tuesday", Desc: "n.星期二"},
	Word{Name: "Wednesday", Desc: "n.星期三"},
	Word{Name: "Thursday", Desc: "星期四"},
	Word{Name: "keep", Desc: "v.(身体健康)处于(状况)"},
	Word{Name: "spend", Desc: "v.度过"},
	Word{Name: "weekend", Desc: "n.周末"},
	Word{Name: "Friday", Desc: "n.星期五"},
	Word{Name: "Saturday", Desc: "n.星期六"},
	Word{Name: "Sunday", Desc: "n.星期日"},
	Word{Name: "country", Desc: "n.乡村"},
	Word{Name: "lucky", Desc: "a.幸运的"},
	Word{Name: "church", Desc: "n.教堂"},
	Word{Name: "dairy", Desc: "n.乳品店"},
	Word{Name: "baker", Desc: "n.面包师傅"},
	Word{Name: "grocer", Desc: "n.食品杂货商"},
	Word{Name: "year", Desc: "n.年"},
	Word{Name: "race", Desc: "n.比赛"},
	Word{Name: "town", Desc: "n.城镇"},
	Word{Name: "crowd", Desc: "n.人群"},
	Word{Name: "stand", Desc: "v.站立"},
	Word{Name: "exciting", Desc: "a.使人激动的"},
	Word{Name: "just", Desc: "ad.正好，恰好"},
	Word{Name: "finish", Desc: "n.结尾，结束"},
	Word{Name: "winner", Desc: "n.获胜者"},
	Word{Name: "behind", Desc: "prop.在…之后"},
	Word{Name: "way", Desc: "n.路途"},
	Word{Name: "awful", Desc: "ad.让人讨厌的，坏的"},
	Word{Name: "telephone", Desc: "v.& n.打电话"},
	Word{Name: "time", Desc: "n.次(数)"},
	Word{Name: "answer", Desc: "v.接(电话)"},
	Word{Name: "last", Desc: "a.最后的，前一次的"},
	Word{Name: "phone", Desc: "n.电话(=telephone)"},
	Word{Name: "again", Desc: "ad.又一次地"},
	Word{Name: "say", Desc: "v.说"},
	Word{Name: "week", Desc: "n.周"},
	Word{Name: "London", Desc: "n.伦敦"},
	Word{Name: "suddenly", Desc: "ad.突然地"},
	Word{Name: "bus stop", Desc: "公共汽车车站"},
	Word{Name: "smile", Desc: "v.微笑"},
	Word{Name: "pleasantly", Desc: "ad.愉快地"},
	Word{Name: "understand", Desc: "v.懂，明白"},
	Word{Name: "speak", Desc: "v.讲，说"},
	Word{Name: "hand", Desc: "n.手"},
	Word{Name: "pocket", Desc: "n.衣袋"},
	Word{Name: "phrasebook", Desc: "n.短语手册，常用语手册"},
	Word{Name: "phrase", Desc: "n.短语"},
	Word{Name: "slowly", Desc: "ad.缓慢地"},
	Word{Name: "hurriedly", Desc: "ad.匆忙地"},
	Word{Name: "cut", Desc: "v.割，切"},
	Word{Name: "thirstily", Desc: "ad.口渴地"},
	Word{Name: "go", Desc: "v.走"},
	Word{Name: "greet", Desc: "v.问候，找招呼"},
	Word{Name: "ago", Desc: "ad.以前"},
	Word{Name: "buy", Desc: "v.买"},
	Word{Name: "pair", Desc: "n.双，对"},
	Word{Name: "fashion", Desc: "n.(服装的)流行式样"},
	Word{Name: "uncomfortable", Desc: "a.不舒服的"},
	Word{Name: "wear", Desc: "v.穿着"},
	Word{Name: "appointment", Desc: "n.约会，预约"},
	Word{Name: "urgent", Desc: "a.紧急的，急迫的"},
	Word{Name: "till", Desc: "prep.直到…为止"},
	Word{Name: "shopping", Desc: "n.购物"},
	Word{Name: "list", Desc: "n.单子"},
	Word{Name: "vegetable", Desc: "n.蔬菜"},
	Word{Name: "need", Desc: "v.需要"},
	Word{Name: "hope", Desc: "v.希望"},
	Word{Name: "thing", Desc: "n.事情"},
	Word{Name: "money", Desc: "n.钱"},
	Word{Name: "groceries", Desc: "n.食品杂货"},
	Word{Name: "fruit", Desc: "n.水果"},
	Word{Name: "stationery", Desc: "n.文具"},
	Word{Name: "newsagent", Desc: "n.报刊零售人"},
	Word{Name: "chemist", Desc: "n.化剂师，化学家"},
	Word{Name: "bath", Desc: "n.洗澡"},
	Word{Name: "nearly", Desc: "ad.几乎，将近"},
	Word{Name: "ready", Desc: "a.准备好的，完好的"},
	Word{Name: "dinner", Desc: "n.正餐，晚餐"},
	Word{Name: "restaurant", Desc: "n.饭馆，餐馆"},
	Word{Name: "roast", Desc: "a.烤的"},
	Word{Name: "breakfast", Desc: "n.早饭"},
	Word{Name: "haircut", Desc: "n.理发"},
	Word{Name: "party", Desc: "n.聚会"},
	Word{Name: "holiday", Desc: "n.假日"},
	Word{Name: "mess", Desc: "n.杂乱，凌乱"},
	Word{Name: "pack", Desc: "v.包装，打包，装箱"},
	Word{Name: "suitcase", Desc: "n.手提箱"},
	Word{Name: "leave", Desc: "v.离开"},
	Word{Name: "already", Desc: "ad.已经"},
	Word{Name: "Paris", Desc: "n.巴黎"},
	Word{Name: "cinema", Desc: "n.电影院"},
	Word{Name: "film", Desc: "n.电影"},
	Word{Name: "beautiful", Desc: "a.漂亮的"},
	Word{Name: "city", Desc: "n.城市"},
	Word{Name: "never", Desc: "ad.从来没有"},
	Word{Name: "ever", Desc: "ad.在任何时候"},
	Word{Name: "attendant", Desc: "n.接待员"},
	Word{Name: "bring", Desc: "v.带来，送来"},
	Word{Name: "garage", Desc: "n.车库，汽车修理厂"},
	Word{Name: "crash", Desc: "n.碰撞"},
	Word{Name: "lamp-post", Desc: "灯杆"},
	Word{Name: "repair", Desc: "v.修理"},
	Word{Name: "try", Desc: "v.努力，设法"},
	Word{Name: "believe", Desc: "v.相信，认为"},
	Word{Name: "may", Desc: "(用于请求许可)可以"},
	Word{Name: "how long", Desc: "多长"},
	Word{Name: "since", Desc: "prep.自从"},
	Word{Name: "why", Desc: "ad.为什么"},
	Word{Name: "sell", Desc: "v.卖，出售"},
	Word{Name: "because", Desc: "conj.因为"},
	Word{Name: "retire", Desc: "v.退休"},
	Word{Name: "cost", Desc: "v.花费"},
	Word{Name: "pound", Desc: "n.英镑"},
	Word{Name: "worth", Desc: "prep.值…钱"},
	Word{Name: "penny", Desc: "n.便士"},
	Word{Name: "still", Desc: "ad.还，仍旧"},
	Word{Name: "move", Desc: "v.搬家"},
	Word{Name: "miss", Desc: "v.想念，思念"},
	Word{Name: "neighbour", Desc: "n.邻居"},
	Word{Name: "person", Desc: "n.人"},
	Word{Name: "people", Desc: "n.人们"},
	Word{Name: "poor", Desc: "a.可怜的"},
	Word{Name: "pilot", Desc: "n.飞行员"},
	Word{Name: "return", Desc: "v.返回"},
	Word{Name: "New York", Desc: "n.纽约"},
	Word{Name: "Tokyo", Desc: "n.东京"},
	Word{Name: "Madrid", Desc: "n.马德里"},
	Word{Name: "fly", Desc: "v.飞行"},
	Word{Name: "Athens", Desc: "n.雅典"},
	Word{Name: "Berlin", Desc: "n.柏林"},
	Word{Name: "Bombay", Desc: "n.孟买"},
	Word{Name: "Geneva", Desc: "n.日内瓦"},
	Word{Name: "Moscow", Desc: "n.莫斯科"},
	Word{Name: "Rome", Desc: "n.罗马"},
	Word{Name: "Seoul", Desc: "n.汉城"},
	Word{Name: "Stockholm", Desc: "n.斯德哥尔摩"},
	Word{Name: "Sydney", Desc: "n.悉尼"},
	Word{Name: "return", Desc: "n.往返"},
	Word{Name: "train", Desc: "n.火车"},
	Word{Name: "platform", Desc: "n.站台"},
	Word{Name: "plenty", Desc: "n.大量"},
	Word{Name: "bar", Desc: "n.酒吧"},
	Word{Name: "station", Desc: "n.车站，火车站"},
	Word{Name: "catch", Desc: "v.赶上"},
	Word{Name: "miss", Desc: "v.错过"},
	Word{Name: "leave", Desc: "v.遗留"},
	Word{Name: "describe", Desc: "v.描述"},
	Word{Name: "zip", Desc: "n.拉链"},
	Word{Name: "label", Desc: "n.标签"},
	Word{Name: "handle", Desc: "n.提手，把手"},
	Word{Name: "address", Desc: "n.地址"},
	Word{Name: "pence", Desc: "n.penny的复数形式"},
	Word{Name: "belong", Desc: "v.属于"},
	Word{Name: "ow", Desc: "int.哎哟"},
	Word{Name: "slip", Desc: "v.滑倒，滑了一脚"},
	Word{Name: "fall", Desc: "v.落下，跌倒"},
	Word{Name: "downstairs", Desc: "ad.下楼"},
	Word{Name: "hurt", Desc: "v.伤，伤害，疼痛"},
	Word{Name: "back", Desc: "n.背"},
	Word{Name: "stand up", Desc: "起立，站起来"},
	Word{Name: "help", Desc: "v.帮助"},
	Word{Name: "at once", Desc: "立即"},
	Word{Name: "sure", Desc: "a.一定的，确信的"},
	Word{Name: "X-ray", Desc: "n.X光透视"},
	Word{Name: "Scotland", Desc: "n.苏格兰(英国)"},
	Word{Name: "card", Desc: "n.明信片"},
	Word{Name: "youth", Desc: "n.青年"},
	Word{Name: "hostel", Desc: "n.招待所，旅馆"},
	Word{Name: "association", Desc: "n.协会"},
	Word{Name: "soon", Desc: "ad.不久"},
	Word{Name: "write", Desc: "v.写"},
	Word{Name: "exam", Desc: "n.考试"},
	Word{Name: "pass", Desc: "v.及格，通过"},
	Word{Name: "mathematics", Desc: "n.数学"},
	Word{Name: "question", Desc: "n.问题"},
	Word{Name: "easy", Desc: "a.容易的"},
	Word{Name: "enough", Desc: "ad.足够地"},
	Word{Name: "paper", Desc: "n.考卷"},
	Word{Name: "fail", Desc: "v.未及格，失败"},
	Word{Name: "answer", Desc: "v.回答"},
	Word{Name: "mark", Desc: "n.分数"},
	Word{Name: "rest", Desc: "n.其他的东西"},
	Word{Name: "difficult", Desc: "a.困难的"},
	Word{Name: "hate", Desc: "v.讨厌"},
	Word{Name: "low", Desc: "a.低的"},
	Word{Name: "cheer", Desc: "v.振作，振奋"},
	Word{Name: "guy", Desc: "n.家伙，人"},
	Word{Name: "top", Desc: "n.上方，顶部"},
	Word{Name: "clever", Desc: "a.聪明的"},
	Word{Name: "stupid", Desc: "a.笨的"},
	Word{Name: "cheap", Desc: "a.便宜的"},
	Word{Name: "expensive", Desc: "a.贵的"},
	Word{Name: "fresh", Desc: "a.新鲜的"},
	Word{Name: "stale", Desc: "a.变馊的"},
	Word{Name: "low", Desc: "a.低的，矮的"},
	Word{Name: "loud", Desc: "a.大声的"},
	Word{Name: "high", Desc: "a.高的"},
	Word{Name: "hard", Desc: "a.硬的"},
	Word{Name: "sweet", Desc: "a.甜的"},
	Word{Name: "soft", Desc: "a.软的"},
	Word{Name: "sour", Desc: "a.酸的"},
	Word{Name: "spell", Desc: "v.拼写"},
	Word{Name: "intelligent", Desc: "a.聪明的，有智慧的"},
	Word{Name: "mistake", Desc: "n.错误"},
	Word{Name: "present", Desc: "n.礼物"},
	Word{Name: "dictionary", Desc: "n.词典"},
	Word{Name: "carry", Desc: "v.携带"},
	Word{Name: "correct", Desc: "v.改正，纠正"},
	Word{Name: "madam", Desc: "n.夫人，女士(对妇女的尊称)"},
	Word{Name: "as well", Desc: "同样"},
	Word{Name: "suit", Desc: "v.适于"},
	Word{Name: "pretty", Desc: "a.漂亮的"},
	Word{Name: "idea", Desc: "n.主意"},
	Word{Name: "a little", Desc: "少许(用于不可数名词之前)"},
	Word{Name: "teaspoonful", Desc: "n.一满茶匙"},
	Word{Name: "less", Desc: "a.(little的比较级)校少的，更小的"},
	Word{Name: "a few", Desc: "几个(用于可数名词之前)"},
	Word{Name: "pity", Desc: "n.遗憾"},
	Word{Name: "instead", Desc: "ad.代替"},
	Word{Name: "advice", Desc: "n.建议，忠告"},
	Word{Name: "most", Desc: "a.(many,much的最高级)最多的"},
	Word{Name: "least", Desc: "a.(little的最高级)最小的，最少的"},
	Word{Name: "best", Desc: "a.(good的最高级)最好的"},
	Word{Name: "worse", Desc: "a.(bad的比较级)更坏的"},
	Word{Name: "worst", Desc: "a.(bad的最高级)最坏的"},
	Word{Name: "model", Desc: "n.型号，式样"},
	Word{Name: "afford", Desc: "v.付得起(钱)"},
	Word{Name: "deposit", Desc: "n.预付定金"},
	Word{Name: "instalment", Desc: "n.分期付款"},
	Word{Name: "price", Desc: "n.价格"},
	Word{Name: "millionaire", Desc: "n.百万富翁"},
	Word{Name: "conductor", Desc: "n.售票员"},
	Word{Name: "fare", Desc: "n.车费，车票"},
	Word{Name: "change", Desc: "v.兑换(钱)"},
	Word{Name: "note", Desc: "n.纸币"},
	Word{Name: "passenger", Desc: "n.乘客"},
	Word{Name: "none", Desc: "pron.没有任何东西"},
	Word{Name: "neither", Desc: "ad.也不"},
	Word{Name: "get off", Desc: "下车"},
	Word{Name: "tramp", Desc: "n.流浪汉"},
	Word{Name: "except", Desc: "prep.除…外"},
	Word{Name: "anyone", Desc: "pron.(用于疑问句，否定式)任何人"},
	Word{Name: "knock", Desc: "v.敲，打"},
	Word{Name: "everything", Desc: "pron.一切事物"},
	Word{Name: "quiet", Desc: "a.宁静的，安静的"},
	Word{Name: "impossible", Desc: "a.不可能的"},
	Word{Name: "invite", Desc: "v.邀请"},
	Word{Name: "anything", Desc: "pron.任何东西"},
	Word{Name: "nothing", Desc: "pron.什么也没有"},
	Word{Name: "lemonade", Desc: "n.柠檬水"},
	Word{Name: "joke", Desc: "v.开玩笑"},
	Word{Name: "asleep", Desc: "a.睡觉，睡着(用作表语)"},
	Word{Name: "glasses", Desc: "n.眼镜"},
	Word{Name: "dining room", Desc: "饭厅"},
	Word{Name: "coin", Desc: "n.硬币"},
	Word{Name: "mouth", Desc: "n.嘴"},
	Word{Name: "swallow", Desc: "v.吞下"},
	Word{Name: "later", Desc: "ad.后来"},
	Word{Name: "toilet", Desc: "n.厕所，盥洗室"},
	Word{Name: "ring", Desc: "v.响"},
	Word{Name: "story", Desc: "n.故事"},
	Word{Name: "happen", Desc: "v.发生"},
	Word{Name: "thief", Desc: "n.贼"},
	Word{Name: "enter", Desc: "v.进入"},
	Word{Name: "dark", Desc: "a.黑暗的"},
	Word{Name: "torch", Desc: "n.手电筒"},
	Word{Name: "voice", Desc: "n.(说话的)声音"},
	Word{Name: "parrot", Desc: "n.鹦鹉"},
	Word{Name: "exercise book", Desc: "练习本"},
	Word{Name: "customer", Desc: "n.顾客"},
	Word{Name: "forget", Desc: "v.忘记"},
	Word{Name: "manager", Desc: "n.经理"},
	Word{Name: "serve", Desc: "v.照应，服务，接待"},
	Word{Name: "counter", Desc: "n.柜台"},
	Word{Name: "recognize", Desc: "v.认识"},
	Word{Name: "road", Desc: "n.路"},
	Word{Name: "during", Desc: "prep.在…期间"},
	Word{Name: "trip", Desc: "n.旅行"},
	Word{Name: "travel", Desc: "v.旅行"},
	Word{Name: "offer", Desc: "v.提供"},
	Word{Name: "job", Desc: "n.工作"},
	Word{Name: "guess", Desc: "v.猜"},
	Word{Name: "grow", Desc: "v.长，让…生长"},
	Word{Name: "beard", Desc: "n.(下巴上的)胡子，络腮胡子"},
	Word{Name: "kitten", Desc: "n.小猫"},
	Word{Name: "water", Desc: "v.浇水"},
	Word{Name: "terribly", Desc: "ad.非常"},
	Word{Name: "dry", Desc: "a.干燥的，干的"},
	Word{Name: "nuisance", Desc: "n.讨厌的东西或人"},
	Word{Name: "mean", Desc: "v.意味着，意思是"},
	Word{Name: "surprise", Desc: "n.惊奇，意外的事"},
	Word{Name: "immediately", Desc: "ad.立即地"},
	Word{Name: "famous", Desc: "a.著名的"},
	Word{Name: "actress", Desc: "n.女演员"},
	Word{Name: "at least", Desc: "至少"},
	Word{Name: "actor", Desc: "n.男演员"},
	Word{Name: "read", Desc: "v.通过阅读得知"},
	Word{Name: "wave", Desc: "v.招手"},
	Word{Name: "track", Desc: "n.跑道"},
	Word{Name: "mile", Desc: "n.英里"},
	Word{Name: "overtake", Desc: "v.从后面超越，超车"},
	Word{Name: "speed", Desc: "限速"},
	Word{Name: "dream", Desc: "v.做梦，思想不集中"},
	Word{Name: "sign", Desc: "n.标记，牌子"},
	Word{Name: "driving licence", Desc: "驾驶执照"},
	Word{Name: "charge", Desc: "v.罚款"},
	Word{Name: "darling", Desc: "n.亲爱的(用作表示称呼)"},
	Word{Name: "Egypt", Desc: "n.埃及"},
	Word{Name: "abroad", Desc: "ad.国外"},
	Word{Name: "worry", Desc: "v.担忧"},
	Word{Name: "reporter", Desc: "n.记者"},
	Word{Name: "sensational", Desc: "a.爆炸性的，耸人听闻的"},
	Word{Name: "mink coat", Desc: "貂皮大衣"},
	Word{Name: "future", Desc: "n.未来的"},
	Word{Name: "get married", Desc: "结婚"},
	Word{Name: "hotel", Desc: "n.饭店"},
	Word{Name: "latest", Desc: "a.最新的"},
	Word{Name: "introduce", Desc: "v.介绍"},
	Word{Name: "football", Desc: "n.足球"},
	Word{Name: "pool", Desc: "n.赌注"},
	Word{Name: "win", Desc: "v.赢"},
	Word{Name: "world", Desc: "n.世界"},
	Word{Name: "poor", Desc: "a.贫穷的"},
	Word{Name: "depend", Desc: "v.依靠(on)"},
	Word{Name: "extra", Desc: "a.额外的"},
	Word{Name: "engineer", Desc: "n.工程师"},
	Word{Name: "overseas", Desc: "a.海外的，国外的"},
	Word{Name: "engineering", Desc: "n.工程"},
	Word{Name: "company", Desc: "n.公司"},
	Word{Name: "line", Desc: "n.线路"},
	Word{Name: "excited", Desc: "a.兴奋的"},
	Word{Name: "get on", Desc: "登上"},
	Word{Name: "middle-aged", Desc: "a.中年的"},
	Word{Name: "opposite", Desc: "prep.在…对面"},
	Word{Name: "curiously", Desc: "ad.好奇地"},
	Word{Name: "funny", Desc: "a.可笑的，滑稽的"},
	Word{Name: "powder", Desc: "n.香粉"},
	Word{Name: "compact", Desc: "n.带镜的化妆盒"},
	Word{Name: "kindly", Desc: "ad.和蔼地"},
	Word{Name: "ugly", Desc: "a.丑陋的"},
	Word{Name: "amused", Desc: "a.有趣的"},
	Word{Name: "smile", Desc: "v.微笑"},
	Word{Name: "embarrassed", Desc: "a.尴尬的，窘迫的"},
	Word{Name: "worried", Desc: "a.担心，担忧"},
	Word{Name: "regularly", Desc: "ad.经常地，定期地"},
	Word{Name: "surround", Desc: "v.包围"},
	Word{Name: "wood", Desc: "n.树林"},
	Word{Name: "beauty spot", Desc: "风景点"},
	Word{Name: "hundred", Desc: "n.百"},
	Word{Name: "city", Desc: "n.城市"},
	Word{Name: "through", Desc: "prep.穿过"},
	Word{Name: "visitor", Desc: "n.参观者，游客，来访者"},
	Word{Name: "tidy", Desc: "a.整齐的"},
	Word{Name: "litter", Desc: "n.杂乱的东西"},
	Word{Name: "litter basket", Desc: "废物筐"},
	Word{Name: "place", Desc: "v.放"},
	Word{Name: "throw", Desc: "v.扔，抛"},
	Word{Name: "rubbish", Desc: "n.垃圾"},
	Word{Name: "count", Desc: "v.数，点"},
	Word{Name: "cover", Desc: "v.覆盖"},
	Word{Name: "piece", Desc: "n.碎片"},
	Word{Name: "tyre", Desc: "n.轮胎"},
	Word{Name: "rusty", Desc: "a.生锈的"},
	Word{Name: "among", Desc: "prep.在…之间"},
	Word{Name: "prosecute", Desc: "v.依法处置"},
}