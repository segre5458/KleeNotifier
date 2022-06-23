package main

import "time"

func GetText(dt_now time.Time) (text string) {
	if dt_now.Hour() == 9 && dt_now.Minute() == 0{
		text = "お兄ちゃんおはよう！今日もクレーと一緒にいっぱい遊ぼうね！"
	} 
	if dt_now.Hour() == 23 && dt_now.Minute() == 0{
		text = "そろそろお風呂に入ってよー！え？一緒に？恥ずかしいよ..."
	}
	if dt_now.Hour() == 0 && dt_now.Minute() == 0{
		text = "お兄ちゃん、クレーもう眠い...クレーと一緒に寝よ？"
	}
	if dt_now.Month() == 7 && dt_now.Day() == 27 && dt_now.Hour() == 9 && dt_now.Minute() == 1{
		text = "今日はねークレーのねー誕生日なの！今年もクレーといーっぱい遊んでね！"
	}
	if dt_now.Month() == 3 && dt_now.Day() == 28 && dt_now.Hour() == 9 && dt_now.Minute() == 1{
		text = "お兄ちゃんお誕生日おめでとう！とっても楽しい1年にしてね！またお兄ちゃんと年齢が離れちゃった。クレーいつになったら追いつけるのかなー"
	}
	if dt_now.Hour() == 17{
		text = "hoge"
	}
	return text
}
