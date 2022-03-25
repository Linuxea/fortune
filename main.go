package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func main() {
	fmt.Println("workout 工作中")

	whoami()
	workout()

}

func whoami() {
	b, err := exec.Command("whoami").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("whoami", string(b))
}

func fortune() string {
	b, err2 := exec.Command("/usr/games/fortune").Output()
	if err2 != nil {
		fmt.Println(err2)
		return "fortune error"
	}

	return string(b)
}

func workout() {

	// notifyUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=4a7bbc6e-f269-480a-b14f-a74f25a8a936&debug=1"
	notifyUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=4a7bbc6e-f269-480a-b14f-a74f25a8a936"

	// 周末不上班
	w := time.Now().Local().Weekday()
	if w == time.Sunday || w == time.Saturday {
		fmt.Println("周末不上班")
		return
	}

	now := time.Now()

	// 0.太早了不通知
	if now.Hour() < 15 {
		return
	}

	// 1.下班提醒
	if now.Hour() == 18 && now.Minute() == 30 {
		tip := "辛苦啦！下班时间已经到了。请整理好桌上东西并记得打卡"
		_, _ = http.Post(notifyUrl, "application/json", bytes.NewBuffer([]byte(buildContext(tip))))

		// image
		s := buildImage("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")
		_, _ = http.Post(notifyUrl, "application/json", bytes.NewBuffer([]byte(s)))
		return
	}

	zero := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	workoutTime := zero.Add(18 * time.Hour).Add(30 * time.Minute)

	totalSecs := int64(workoutTime.Sub(now).Seconds())
	hours := totalSecs / 3600
	minutes := (totalSecs % 3600) / 60
	seconds := totalSecs % 60

	if hours < 0 || minutes < 0 || seconds < 0 {
		return
	}

	// 2.下班时长提醒
	cont := fmt.Sprintf("距离下班时间还有: %02d:%02d:%02d, \n\n%s\n\n%s", hours, minutes, seconds, strings.TrimSpace(luxun()), fortune())
	_, httpErr := http.Post(notifyUrl, "application/json", bytes.NewBuffer([]byte(buildContext(cont))))
	if httpErr != nil {
		fmt.Println("http 异常", httpErr.Error())
	}

	// 3.狗命提醒
	if now.Minute() == 30 {
		_, httpErr := http.Post(notifyUrl, "application/json", bytes.NewBuffer([]byte(buildContext("起来走走吧，命是自己的 钱是公司的"))))
		if httpErr != nil {
			fmt.Println("http 异常", httpErr.Error())
		}
	}

	// 4.瑞幸提醒
	if (w == time.Wednesday || w == time.Monday) && now.Minute() == 30 && now.Hour() == 15 {
		_, httpErr := http.Post(notifyUrl, "application/json", bytes.NewBuffer([]byte(buildContext("来杯瑞幸？能饮一杯无。yyds? "))))
		if httpErr != nil {
			fmt.Println("http 异常", httpErr.Error())
		}
	}

	// 5.一半时间过去提醒
	if now.Hour() == 15 && now.Minute() == 0 {
		_, httpErr := http.Post(notifyUrl, "application/json", bytes.NewBuffer([]byte(buildContext("一半时间过去了！加油盆友们"))))
		if httpErr != nil {
			fmt.Println("http 异常", httpErr.Error())
		}
	}

}

func buildContext(tip string) string {
	return `
{
	"msgtype": "text",
	"text": {
	"content": ` + "\"" + tip + "\"" + `
	}
}
`
}

func buildImage(imageBase64 string) string {

	return `
	{
		"msgtype": "image",
		"image": {
			"base64": ` + "\"" + imageBase64 + "\"" + `,
			"md5": ` + "\"" + fmt.Sprintf("%s", "f0411b8c767967eaca9cd58272fefeeb") + "\"" + `
		}
	}
	
	`
}

func luxun() string {

	content := `
	改造自己，总比禁止别人来得难。——鲁迅

幼稚对于老成，有如孩子对于老人，决没有什么耻辱的，作品也一样，起初幼稚，不算耻辱的。

不在沉默中爆发，就在沉默中灭亡。

我们中国人对于不是自己的东西，或者将不为自己所有的东西，总要破坏了才快活的。

横眉冷对千夫指，俯首甘为孺子牛

凡是总须研究，才会明白。

愈艰难，就愈要做。改革，是向来没有一帆风顺的。

智识太多，不是心活，就是心软。心活就会胡思乱想，心软就不肯下辣子手……所以智识非铲除不可。

节省时间，也就是使一个人的有限的生命更加有效，而也即等于延长了人的生命。(鲁迅)

那里有天才，我是把别人喝咖啡的工夫都用在工作上的。——鲁迅

寄意寒星荃不察，我以我血荐轩辕。(鲁迅?自题小像)

杀了“现在”，也便杀了“将来”。----将来是子孙的时代。——鲁迅

天才并不是自生自长在深林荒野里的怪物，是由可以使天才生长的民众产生、长育出来的，所以没有这种民众，就没有天才。——鲁迅

改造自己，总比禁止别人来的难。

巨大的建筑，总是由一木一石叠起来的，我们何妨做做这一木一石呢?我时常做些零碎事，就是为此。——鲁迅

激烈得快的，也平和的快，甚至于也颓废的快。

要竭力将可有可无的字、句、段删去，毫不可惜。(鲁迅)

横眉冷对千夫指，俯首干为孺子牛。

忍看朋辈成新鬼，怒向刀丛觅小诗。(鲁迅。《为了忘却的纪念》)

事实是毫无情面的东西，它能将空言打得粉碎。

必须敢于正视，这才可望敢想、敢说、敢做、敢当。

我每看运动会时，常常这样想：优胜者固然可敬，但那虽然落后而仍非跑至终点的竞技者，和见了这样的竞技者而肃然不笑的看客，乃正是中国将来之脊梁。

谈之类，是谈不久，也谈不出什麽来的，它始终被事实的镜子照出原形，拖出尾巴而去。——鲁迅

空谈之类，是谈不久，也谈不出什麽来的，它始终被事实的镜子照出原形，拖出尾巴而去。——鲁迅

游戏是儿童最正当的行为，玩具是儿童的天使。

说过的话不算数，是中国人的大毛病。

时间就像海绵里的水，只要愿挤，总还是有的。

人生得一知己足矣，斯世当以同怀视之。——鲁迅

写小说，说到底，就是写人物。小说艺术的精髓就是创造人物的艺术。(鲁迅)

敌人是不足惧的，最可怕的是自己营垒里的蛀虫，许多事情都败在他们手里。

无情未必真豪杰，怜子如何不丈夫

与名流者谈，对于他之所讲，当装作偶有不懂之处。太不懂被看轻，太懂了被厌恶。偶有不懂之处，彼此最为合宜。

希望是附丽于存在的，有存在，便有希望，有希望，便是光明。——鲁迅

从喷泉里出来的都是水，从血管里出来的都是血。(鲁迅)

人类的悲欢并不相通……

我好象一只牛，吃的是草，挤出的是奶、血。——鲁迅

倘只看书，便变成书橱。

明言着轻蔑什么人，并不是十足的轻蔑。惟沉默是最高的轻蔑-------最高的轻蔑是无言，而且连眼珠也不转过去。

我好象是一只牛，吃的是草，挤出的是牛奶。

唯有民族魂是值得宝贵的，唯有它发扬起来，中国才有真进步。

中国的有一些士大夫，总爱无中生有，移花接木地造出故事来，他们不但歌颂生平，还粉饰黑暗。

勇者愤怒，抽刃向更强者;怯者愤怒，却抽刃向更弱者。不可救药的民族中，一定有许多英雄，专向孩子们瞪眼。这些孱头们。

从来如此，便对吗?

我们和朋友在一起，可以脱掉衣服，但上阵要穿甲。

一件事，无论大小，倘无恒心，是很不好的。而看一切太难，固然能使人无成，但若看得太容易，也能使事情无结果。

当我沉默的时候，我觉得充实;我将开口，同时感到空虚。

哪里有天才，我是把别人喝咖啡的工夫都用在了工作上了。

惟有民魂是值得宝贵的，惟有他发扬起来，中国才有真进步。——鲁迅

其实地上本没有路，走的人多了，也便成了路。

寄意寒星荃不察，我以我血荐轩辕。

度尽劫波兄弟在，相逢一笑泯恩仇。(鲁迅题三义塔)

时间就是性命。无端的空耗别人的时间，其实是无异于谋财害命。

单是说不行，要紧的是做。——鲁迅

孩子是要别人教的，毛病是要别人医的，即使自己是教员或医生。但做人处事的法子，却恐怕要自己斟酌，许多人开来的良方，往往不过是废纸。——鲁迅

史家之绝唱，无韵之《离骚》。(鲁迅评《史记》)

中国一向就少有失败的英雄，少有韧性的反抗，少有敢单身鏖战的武人，少有敢抚哭叛徒的吊客;见胜兆则纷纷聚集，见败兆则纷纷逃亡。

心事浩茫连广宇，于无声处听惊雷。--鲁迅

血沃中原肥劲草，寒凝大地发春华。--鲁迅

做一件事，无论大小，倘无恒心，是很不好的。而看一切太难，固然能使人无成，但若看得太容易，也能使事情无结果。

哈儿狗往往比它的主人更严厉。

有一分热，发一分光。

无情未必真豪杰，怜子如何不丈夫

改造自己，总比禁止别人来得难。——鲁迅

我们目下的当务之急是：一要生存，二要温饱，三要发展。

岂有豪情似旧时，花开花落两由之。(鲁迅?悼杨铨)

我自爱我的野草，但我憎恶这以野草作装饰的地面。

横眉冷对千夫指，俯首甘为孺子牛。(鲁迅?自嘲)

宁可与敌人明打，不欲受同人暗算。

贪安稳就没有自由，要自由就要历些危险。只有这两条路。

不耻最后”。即使慢，驰而不息，纵会落后，纵令失败，但一定可以达到他所向的目标。

曾经阔气的要复古，正在阔气的要保持现状，未曾阔气的要革新，大抵如此，大抵!

没有冲破一切传统思想和手法的闯将，中国不会有真的新文艺的。文艺是国民精神所发的火光，同时也是引导国民精神前途的灯光。(鲁迅)

真的勇士敢于直面惨淡的人生，敢于正视淋漓的鲜血。

其实地上本没有路，走的人多了，也便成了路。”

悲剧将人生的有价值的东西毁灭给人看，喜剧将那无价值的撕破给人看。(鲁迅)

搞鬼有术，也有效，然而有限，所以以此成大事者，古来无有。

时间就是生命，无端地空耗别人的时间，无异于谋财害命。

沉着、勇猛，有辨别，不自私。

走自己的路，让别人去说

巨大的建筑，总是一木一石叠起来，我们何尝做做这一木一石呢?我时常做些零碎事，就是为此。

鲁迅的经典语录

假使做事要面面顾到，那就什么事都不能做了。

愿中国青年都摆脱冷气，只是向上走，不必听自暴自弃者的话。

无情未必真豪杰，怜子如何不丈夫。(鲁迅)

世上本没有路，走的人多了，也便成了路

心事浩茫连广宇，于无声处听惊雷。(鲁迅?无题)

做一件事，无论大小，倘无恒心，是很不好的。

人类总不会寂寞，以为生命是进步的，是天生的。

我们自古以来，就有埋头苦干的人，有拼命硬干的人，有为民请命的人，有舍身求法的人……

上人生的旅途罢。前途很远，也很暗。然而不要怕。不怕的人的面前才有路。——鲁迅

只要能培一朵花，就不妨做做会朽的腐草。

大的建筑，总是由一木一石叠起来的，我们何妨做做这一木一石呢?我时常做些零碎事，就是为此。——鲁迅

不满是向上的车轮。(鲁迅)

倘能生存，我当然仍要学习。

在行进时，也时时有人退伍，有人落荒，有人颓唐，有人叛变，然而只要无碍于进行，则越到后来，这队伍也就越成为纯粹、精锐的队伍了。——鲁迅

自由固不是钱所买到的，但能够为钱而卖掉。——鲁迅

死者倘不埋在活人心中，那就真的死掉了。

血沃中原肥劲草，寒凝大地发春华。(鲁迅_无题)

疑并不是缺点。总是疑，而并不下断语，这才是缺点。——鲁迅

墨写的谎说，决掩不住血写的事实。

其实先驱者本是容易变成绊脚石的。

一个人的价值，应该看他贡献什么，而不应当看他取得什么。——爱因斯坦

　　共同的事业，共同的斗争，可以使人们产生忍受一切的力量。——奥斯特洛夫斯基

　　古之立大事者，不惟有超世之才，亦必有坚忍不拔之志。——苏轼

　　理想的人物不仅要在物质需要的满足上，还要在精神旨趣的满足上得到表现。——黑格尔

　　读一本好书，就如同和一个高尚的人在交谈。——歌德

　　过去一切时代的精华尽在书中。——卡莱尔

　　好的书籍是最贵重的珍宝。——别林斯基

　　读书是易事，思索是难事，但两者缺一，便全无用处。——富兰克林

　　读书是在别人思想的帮助下，建立起自己的思想。——鲁巴金

　　合理安排时间，就等于节约时间。——培根

　　你想成为幸福的人吗？但愿你首先学会吃得起苦。——屠格涅夫

　　抛弃时间的人，时间也抛弃他。——莎士比亚

　　普通人只想到如何度过时间，有才能的人设法利用时间。——叔本华

　　读书破万卷，下笔如有神。——杜甫

　　取得成就时坚持不懈，要比遭到失败时顽强不屈更重要。——拉罗什夫科

　　人的一生是短的，但如果卑劣地过这一生，就太长了。——莎士比亚

　　读书忌死读，死读钻牛角。——叶圣陶

　　不要回避苦恼和困难，挺起身来向它挑战，进而克服它。——池田大作

　　真实是人生的命脉，是一切价值的根基。——德莱塞

　　努力爱一个人。付出，不一定会有收获；不付出，却一定不会有收获，不要奢望出现奇迹。

　　土地是以它的肥沃和收获而被估价的；才能也是土地，不过它生产的不是粮食，而是真理。如果只能滋生瞑想和幻想的话，即使再大的才能也只是砂地或盐池，那上面连小草也长不出来的。——别林斯基

　　新的数学方法和概念，常常比解决数学问题本身更重要。——华罗庚

　　人生的磨难是很多的，所以我们不可对于每一件轻微的伤害都过于敏感。在生活磨难面前，精神上的坚强和无动于衷是我们抵抗罪恶和人生意外的最好武器。——洛克

　　友谊是灵魂的结合，这个结合是可以离异的，这是两个敏感，正直的人之间心照不宣的契约。——伏尔泰

　　人生并不像火车要通过每个站似的经过每一个生活阶段。人生总是直向前行走，从不留下什么。——刘易斯

　　都睡着的时候，一步步艰辛地向上攀爬的。

　　劳动和人，人和劳动，这是所有真理的父母亲。——苏霍姆林斯基

　　伟人之所以伟大，是因为他与别人共处逆境时，别人失去了信心，他却下决心实现自己的目标。

　　笨蛋自以为聪明，聪明人才知道自己是笨蛋。——莎士比亚

　　不要忘本，任何时候，任何事情。

　　不为失败找理由，要为成功找方法。

　　谁不会休息，谁就不会工作。——列宁

　　回避现实的人，未来将更不理想。

　　友谊是一棵可以庇荫的树。——柯尔律治

　　个人的智慧只是有限的。——普劳图斯

　　世界上那些最容易的事情中，拖延时间最不费力。

　　当你感到悲哀痛苦时，最好是去学些什么东西。学习会使你永远立于不败之地。

　　良好的健康状况和高度的身体训练，是有效的脑力劳动的重要条件。——克鲁普斯卡娅

　　出门走好路，出口说好话，出手做好事。

　　爱情只有当它是自由自在时，才会叶茂花繁。认为爱情是某种义务的思想只能置爱情于死地。只消一句话：你应当爱某个人，就足以使你对这个人恨之入骨。——罗素

　　酒杯里竟能蹦出友谊来。——盖伊

　　最甜美的是爱情，最苦涩的也是爱情。——菲·贝利

　　管理的第一目标是使较高工资与较低的劳动成本结合起来。——泰罗

　　科学是到处为家的，不过，在任何不播种的地方，是决不会得到丰收的。——赫尔岑

　　真理惟一可靠的标准就是永远自相符合。——欧文

　　积极的人在每一次忧患中都看到一个机会，而消极的人则在每个机会都看到某种忧患。

　　时间是一切财富中最宝贵的财富。——德奥弗拉斯多

　　伟大的思想能变成巨大的财富。——塞内加

　　温顺的青年人在图书馆里长大，他们相信他们的责任是应当接受西塞罗，洛克，培根发表的意见；他们忘了西塞罗，洛克与培根写这些书的时候，也不过是在图书馆里的青年人。——爱默生

　　成功大易，而获实丰于斯所期，浅人喜焉，而深识者方以为吊。——梁启超

　　说真话不应当是艰难的事情。我所谓真话不是指真理，也不是指正确的话。自己想什麽就讲什麽；自己怎麽想就怎麽说这就是说真话。——巴金

　　坚韧是成功的一大要素，只要在门上敲得够久够大声，终会把人唤醒的。

　　请记得，好朋友的定义是：你混的好，她打心眼里为你开心；你混的不好，她由衷的为你着急。

　　过放荡不羁的生活，容易得像顺水推舟，但是要结识良朋益友，却难如登天。——巴尔扎克

　　躯体总是以惹人厌烦告终。除思想以外，没有什么优美和有意思的东西留下来，因为思想就是生命。——萧伯纳

　　盛年不重来，一日难再晨。及时宜自勉，岁月不待人。——陶渊明

　　将人生投于的赌徒，当他们胆敢妄为的时候，对自己的力量有充分的自信，并且认为大胆的冒险是唯一的形式。——茨威格

　　人的活动如果没有理想的鼓舞，就会变得空虚而渺小。——车尔尼雪夫斯基

　　从不浪费时间的人，没有工夫抱怨时间不够。——杰弗逊

　　没有口水与汗水，就没有成功的泪水。

　　要诚恳，要坦然，要慷慨，要宽容，要有平常心。

　　怠惰是贫穷的制造厂。

　　较高级复杂的劳动，是这样一种劳动力的表现，这种劳动力比较普通的劳动力需要较高的教育费用，它的生产需要花费较多的劳动时间。因此，具有较高的价值。——马克思

　　我读的书愈多，就愈亲近世界，愈明了生活的意义，愈觉得生活的重要。——高尔基

　　要使整个人生都过得舒适、愉快，这是不可能的，因为人类必须具备一种能应付逆境的态度。——卢梭

　　只有把抱怨环境的心情，化为上进的力量，才是成功的保证。——罗曼·罗兰

　　知之者不如好之者，好之者不如乐之者。——孔子

　　勇猛、大胆和坚定的决心能够抵得上武器的精良。——达·芬奇

　　意志是一个强壮的盲人，倚靠在明眼的跛子肩上。——叔本华

　　只有永远躺在泥坑里的人，才不会再掉进坑里。——黑格尔

　　希望的灯一旦熄灭，生活刹那间变成了一片黑暗。——普列姆昌德

　　希望是人生的乳母。——科策布

　　形成天才的决定因素应该是勤奋。——郭沫若

　　学到很多东西的诀窍，就是一下子不要学很多。——洛克

　　自己的鞋子，自己知道紧在哪里。——西班牙

　　我们唯一不会改正的缺点是软弱。——拉罗什福科

　　我这个人走得很慢，但是我从不后退。——亚伯拉罕·林肯

　　勿问成功的秘诀为何，且尽全力做你应该做的事吧。——美华纳

　　学而不思则罔，思而不学则殆。——孔子

　　学问是异常珍贵的东西，从任何源泉吸收都不可耻。——阿卜·日·法拉兹

　　只有在人群中间，才能认识自己。——德国

　　重复别人所说的话，只需要教育；而要挑战别人所说的话，则需要头脑。——玛丽·佩蒂博恩·普尔

　　卓越的人一大优点是：在不利与艰难的遭遇里百折不饶。——贝多芬

　　自己的饭量自己知道。——苏联

　　我们若已接受最坏的，就再没有什么损失。——卡耐基

　　书到用时方恨少、事非经过不知难。——陆游

　　书籍把我们引入最美好的社会，使我们认识各个时代的伟大智者。——史美尔斯

　　熟读唐诗三百首，不会作诗也会吟。——孙洙

　　谁和我一样用功，谁就会和我一样成功。——莫扎特

　　天下之事常成于困约，而败于奢靡。——陆游

　　生命不等于是呼吸，生命是活动。——卢梭

　　伟大的事业，需要决心，能力，组织和责任感。——易卜生

　　唯书籍不朽。——乔特

　　青春是一个短暂的美梦，当你醒来时，它早已消失无踪。——莎士比亚

　　书不仅是生活，而且是现在、过去和未来文化生活的源泉。——库法耶夫

　　生命不可能有两次，但许多人连一次也不善于度过。——吕凯特

　　问渠哪得清如许，为有源头活水来。——朱熹

　　我的努力求学没有得到别的好处，只不过是愈来愈发觉自己的无知。——笛卡儿

　　生活的道路一旦选定，就要勇敢地走到底，决不回头。——左拉

　　奢侈是舒适的，否则就不是奢侈。

　　书籍是朋友，虽然没有热情，但是非常忠实。——雨果

　　伟大的成绩和辛勤劳动是成正比例的，有一分劳动就有一分收获，日积月累，从少到多，奇迹就可以创造出来。——鲁迅

　　非淡泊无以明志，非宁静无以致远。——诸葛亮

　　书痴者文必工，艺痴者技必良。——蒲松龄

　　读过一本好书，像交了一个益友。——藏克家

　　读书以过目成诵为能，最是不济事。——郑板桥

　　书籍是青年人不可分离的生活伴侣和导师。——高尔基

　　鸟欲高飞先振翅，人求上进先读书。——李苦禅

　　立志宜思真品格，读书须尽苦功夫。——阮元

　　不怕读得少，只怕记不牢。——徐特立

　　读不在三更五鼓，功只怕一曝十寒。——郭沫若

　　不读书的人，思想就会停止。——狄德罗

　　聪明在于勤奋，天才在于积累。——华罗庚

　　书籍是屹立在时间的汪洋大海中的灯塔。——惠普尔

　　韬略终须建新国，奋发还得读良书。——郭沫若

　　路漫漫其修道远，吾将上下而求索。——屈原

　　理想的书籍是智慧的钥匙。——列夫·托尔斯泰

　　如果把生活比喻为创作的意境，那么阅读就像阳光。——池莉

　　千里之行，始于足下。——老子

　　处处是创造之地，天天是创造之时，人人是创造之人。——陶行之千教万教教人求真，千学万学学做真人。——陶行之

　　读书也像开矿一样“沙里淘金”。——赵树理

　　时间就像海绵里的水，只要愿挤，总还是有的。——鲁迅

　　经验丰富的人读书用两只眼睛，一只眼睛看到纸面上的话，另一眼睛看到纸的背面。——歌德

　　人的影响短暂而微弱，书的影响则广泛而深远。——普希金

　　一个爱书的人，他必定不致于缺少一个忠实的朋友，一个良好的老师，一个可爱的伴侣，一个温情的安慰者。——巴罗

　　勿以恶小而为之，勿以善小而不为。——陈寿《三国志》

　　阅读的最大理由是想摆脱平庸，早一天就多一份人生的精彩；迟一天就多一天平庸的困扰。——余秋雨

　　读书不要贪多，而是要多加思索，这样的读书使我获益不少。——卢梭

　　饭可以一日不吃，觉可以一日不睡，书不可以一日不读。——毛泽东

　　读书也像开矿一样“沙里淘金”。——赵树理

　　最有希望的成功者，并不是才干出众的人而是那些最善利用每一时机去发掘开拓的人。——苏格拉底

　　告诉你使我达到目标的奥秘吧，我唯一的力量就是我的坚持精神。——巴斯德

　　如果道德败坏了，趣味也必然会堕落。——狄德罗

　　对于一切事物，期望总比绝望好。——歌德

　　失信就是失败。——左拉

　　书有自己的命运，要视读者接受的情况而定。——忒伦提乌斯·摩尔

　　不管怎样的事情，都请安静地愉快吧！这是人生。我们要依样地接受人生，勇敢地大胆地，而且永远地微笑着。——卢森堡

　　发现者，尤其是一个初出茅庐的年轻发现者，需要勇气才能无视他人的冷漠和怀疑，才能坚持自己发现的意志，并把研究继续下去。——贝弗里奇

　　不读书的人，思想就会停止。人都向往知识，一旦知识的渴望在他身上熄灭，他就不再成为人。——南森

　　得道者多助，失道者寡助。——孟子

　　石以砥焉，化钝为利。——刘禹锡

　　我读书越多，书籍就使我和世界越接近，生活对我也变得越加光明和有意义。——高尔基

　　温故而知新，可以为师矣。——孔子

　　天才就是这样，终身努力，便成天才。——门捷列夫

　　只有在苦难中，才能认识自我。——希尔蒂

　　生活这样美好，活它一千辈子吧！——贝多芬

　　值得做的事情必定是困难的。——奥维德

　　意志是独一无二的个体所拥有的以纠正自己的自动性的力量。——劳伦斯

　　人不守信，无异于叫旁人对他失信。——英国

　　书，以是哺育心灵的母乳，启迪智慧的钥匙。——高尔基

　　雨下给富人，也下给穷人，下给义人，也下给不义的人；其实，雨并不公道，因为下落在一个没有公道的世界上。——老舍

　　我想写一出最悲的悲剧，里面充满了无耻的笑声。——老舍

　　婚姻是一座围城，城外的人想进去，城里的人想出来。——钱钟书

　　爱情多半是不成功的，要么苦于终成眷属的厌倦，要么苦于未能终成眷属的悲哀。——钱钟书

　　人总是在接近幸福时倍感幸福，在幸福进行时却患得患失。——张爱玲

　　记忆的梗上，谁不有，两三朵娉婷，披着情绪的花。——林徽因

　　我从没被谁知道，所以也没被谁忘记。在别人的回忆中生活，并不是我的目的。——顾城

　　于千万人之中遇见你所要遇见的人，于千万年之中，时间的无涯的荒野里，没有早一步，也没有晚一步，刚巧赶上了，那也没有别的话可说，唯有轻轻地问一声：“噢，你也在这里吗？”——张爱玲

　　心小了，所有的小事就大了；心大了，所有的大事都小了；看淡世事沧桑，内心安然无恙。——丰子恺

　　爱情里最棒的心态就是：我的一切付出都是一场心甘情愿，我对此绝口不提。你若投桃报李，我会十分感激。你若无动于衷，我也不灰心丧气。直到有一天我不愿再这般爱你，那就让我们一别两宽，各生欢喜。 ——网易云热评《姗姗》

　　如果不幸福，如果不快乐，那就放手吧；如果舍不得、放不下，那就痛苦吧。——柏拉图

　　记住该记住的，忘记该忘记的。改变我能改变的，接受我不能改变的。——塞林格

　　花开半看，酒饮微醉，此中大有佳趣。若是烂漫，便成恶境美。履盈满者，宜思之——洪应明

　　少年时追求激情，成熟后却迷恋平静，在我们寻找，伤害，背离之后，还能一如既往的相信爱情，这是一种勇气。——村上春树

　　不同的青春，同样的迷惘。然而，青春会成长，迷惘会散去，黑夜过后，太阳照常升起——海明威

　　一个人记得事情太多真不幸，知道事情太多也不幸，体会到太多事情也不幸。——沈从文

　　世事茫茫，光阴有限，算来何必奔忙？人生碌碌，竞短论长，却不道荣枯有数，得失难量。——沈复

　　我将于茫茫人海中访我灵魂之伴侣；得之，我幸；不得，我命。——徐志摩

　　我什么都没有忘，只是有些事只适合收藏，不能说，不能想，却也不能忘。——史铁生

　　在一回首间，才忽然发现，原来，我一生的种种努力，不过只为了周遭的人对我满意而已。为了搏得他人的称许与微笑，我战战兢兢地将自己套入所有的模式所有的桎梏。走到途中才忽然发现，我只剩下一副模糊的面目，和一条不能回头的路。——席慕容

　　书籍备而不读，等于废纸。——英国谚语

　　一日不书，百事荒芜。——（唐）李诩

　　知识是为老年准备的最好的食粮。——亚里士多德

　　不知道自己无知，乃是双倍的无知。——柏拉图

　　读书破万卷，下笔如有神。——杜甫

　　知识使智慧者更聪明，使愚昧者更愚蠢。——英国

　　人的知识愈广，人的本身也愈臻完善。——高尔基

　　知是一种快乐，好奇心知识的萌芽。——培根

　　惟一会妨碍我学习的是，我所受到的教育。——爱因斯坦

　　数学是研究抽象结构的理论。——布尔巴基学派

　　谁有历经千辛万苦的意志，谁就能达到任何目的。——米南德

　　科学是使人变得勇敢的最好途径。——布鲁诺

　　运用一分知识，需要十分积累。——伊朗

　　无论你有多少知识，假如不用便是一无所知。——阿拉伯

　　智者的智慧是一种不平常的常识。——拉尔夫

　　知识比金子宝贵，因为金子买不到它。——苏联

　　学者的一天比不学无术的人一生还有价值。——阿拉伯

　　知识是头上的花环，而财产是颈上的枷锁。

　　常识是事物可能性的尺度，由预见和经验组成。——亚美路

　　人生并非游戏，因此我们没有权利随意放弃它。——列夫托尔斯泰

　　学识太广反而憨头憨脑。——罗·伯顿

　　惜时专心苦读是做学问的一个好方法。——蔡尚思

　　我们要像海绵一样吸收有用的知识。——加里宁

　　人如果没有知识，无异于行尸走肉。——托·因哲伦德

　　不尽读天下之书，不能相天下之士。——汤显祖

　　知识能使你增加一双眼睛。——叙利亚

　　科学是“无知”的局部解剖学。——霍姆斯

　　引诱肉体的是金钱和奢望，吸引灵魂的是知识和理智。——伊朗

　　学而不思则罔，思而不学则殆。——孔子

　　知识越少越准确，知识越多，疑惑也就越多。——歌德

　　任何倏忽的灵感事实上不能代替长期的功夫。——罗丹

　　学问对于人们要求最大的紧张和最大的热情。——巴甫洛夫

　　天才就是百分之九十九的汗水加百分之一的灵感。——爱迪生

　　科学发展的终点是哲学，哲学发展的终点是宗教。——杨振宁

　　斗争是掌握本领的学校，挫折是通向真理的桥梁。——歌德

　　不下决心培养思考的人，便失去了生活中的最大乐趣。——爱迪生

　　没有比知识更好的朋友，没有比病魔更坏的敌人。——印度

　　书本上的知识而外，尚须从生活的人生中获得知识。——茅盾

　　真理是严酷的，我喜爱这个严酷，它永不欺骗。——泰戈尔

　　理想的书籍，是智慧的钥匙。——列夫·托尔斯泰
	`

	rand.Seed(time.Now().Unix())
	s := strings.Split(content, "\n")
	for {
		s2 := s[rand.Int31n(int32(len(s)))]

		if s2 == "" {
			continue
		}

		return s2

	}

}
