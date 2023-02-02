package lang

var ja_JP map[string]string = map[string]string{
	"Deposit":             "入金",
	"Withdraw":            "金のうち",
	"Buy product":         "製品を購入",
	"Product settlement":  "製品決済",
	"Deposit gift":        "入金贈与",
	"System deposit":      "システム預金",
	"System deduction":    "システム控除",
	"System freezes":      "システムフリーズ",
	"System unfreeze":     "システム解凍",
	"Invest transfer in":  "投資転送",
	"Invest transfer out": "投資振り出し",
	"Investment income":   "投資収益率",

	"Parameter error":                       "パラメータ エラー",
	"Username cannot be empty":              "ユーザー名を空にすることはできません",
	"Password cannot be empty":              "パスワードを空にすることはできません",
	"The two passwords are inconsistent":    "2 つのパスワードの不整合",
	"Please enter the correct phone number": "正しい携帯電話番号を入力してください",
	"Invitation code cannot be empty":       "招待コードを空にすることはできません",
	"Username already exists":               "ユーザー名は既に存在します",
	"Username does not exist":               "ユーザー名が存在しません",
	"Registration function is closed":       "登録機能オフ中",
	"Wrong invitation code":                 "招待コードエラー",
	"Incorrect username and password":       "ユーザー名のパスワードが正しくありません",
	"User is forbidden to log in":           "ユーザーはログインを禁止しています",

	"User is not logged in":                          "用户未登录",
	"Mobile phone number cannot be empty":            "手机号码不能为空",
	"The delivery address cannot be empty":           "收货地址不能为空",
	"The detailed address cannot be empty":           "详细地址不能为空",
	"Cardholder cannot be empty":                     "持卡人不能为空",
	"Bank card number cannot be empty":               "银行卡号不能为空",
	"Account bank cannot be empty":                   "开户行不能为空",
	"Withdraw password cannot be empty":              "出金密码不能为空",
	"Incorrect withdraw password":                    "出金密码错误",
	"The original password cannot be empty":          "原密码不能为空",
	"The original password is wrong":                 "原密码错误",
	"The original withdraw password cannot be empty": "原出金密码不能为空",

	"The original withdraw password is wrong":                "元の金パスワードエラー",
	"Wrong deposit amount":                                   "入金金額エラー",
	"Wrong amount":                                           "金額エラー",
	"The deposit voucher cannot be empty":                    "預金伝票は空にできません",
	"Record does not exist":                                  "レコードが存在しません",
	"Please deposit after %s":                                "%s の後に入金してください",
	"Please deposit before %s":                               "%s より前に入金してください",
	"Minimum deposit %.2f":                                   "最低入金%.2f",
	"Maximum deposit %.2f":                                   "最大入金%.2f",
	"The start time cannot be greater than the current time": "開始時刻は現在の時刻より大きくすることはできません",
	"The start time cannot be greater than the end time":     "開始時刻は終了時刻より大きくすることはできません",

	"Platform cannot be empty":                       "プラットフォームを空にすることはできません",
	"No app version":                                 "アプリなしバージョン",
	"Receiving address cannot be empty":              "入金アドレスを空にすることはできません",
	"Insufficient account balance":                   "口座残高不足",
	"バンクカードエレオール":                                    "銀行カードエラー",
	"Please withdraw after %s":                       "%s の後に出金してください",
	"Please withdraw before %s":                      "%s より前に出金してください",
	"Minimum withdraw %.2f":                          "最低出金 %.2f",
	"マキシム withdraw %.2f":                             "最大出金 %.2f",
	"System error, please contact the administrator": "システム エラー,管理者に問い合わせてください",
	"The current user is forbidden to withdraw":      "現在のユーザーは出金禁止",
	"Only one shipping address can be added":         "1 つの出荷先住所のみを追加できます",

	"Wrong deposit type":            "入金タイプエラー",
	"Wrong purchase type":           "購入タイプエラー",
	"Order number cannot be empty":  "注文番号を空にすることはできません",
	"Order does not exist":          "注文が存在しません",
	"Platform name cannot be empty": "プラットフォーム名を空にすることはできません",
	"Platform name error":           "プラットフォーム名が正しくありません",
	"System configuration error, please contact the administrator": "システム構成エラー,管理者に連絡してください",
	"Payment account cannot be empty":                              "支払いアカウントを空にすることはできません",
	"Receiving account cannot be empty":                            "入金口座番号を空にすることはできません",
	"Receiving account does not exist":                             "入金口座番号が存在しません",
	"Payment channel cannot be empty":                              "支払チャネルを空にすることはできません",
	"The payment channel does not exist":                           "支払チャネルが存在しません",

	"The payment does not exist":                                "支払いは存在しません",
	"バンクキャントベエプティ":                                              "銀行は空にすることはできません",
	"Unable to bind bank card, please contact customer service": "銀行カードをバインドすることはできません, カスタマーサービスにお問い合わせください",
	"Only %d bank cards can be bound":                           "%d 枚の銀行カードのみをバインドできます",
	"Prohibit to modify bank card":                              "銀行カードの変更禁止",
	"バンクカード does not exist":                                     "銀行カードが存在しません",
	"Please upload a picture within %dM":                        "%dM 内の画像をアップロードしてください",
	"Picture format error":                                      "画像形式が正しくありません",
	"Wrong withdrawal method":                                   "金出し方法が正しくありません",
	"Configuration error":                                       "構成エラー",
	"Prohibited to buy":                                         "購入禁止",
	"Insufficient transferable balance":                         "転送可能な残高が不足しています",

	"Country code cannot be empty":                       "国コードを空にすることはできません",
	"Product does not exist":                             "製品が存在しません",
	"Optional added":                                     "オートプトが追加されました",
	"You can only withdraw %d times per day":             "1 日に 1 回だけ出金 %d 回",
	"The current account has been frozen!":               "現在のアカウントは凍結されています!",
	"Real name cannot be empty":                          "名前を空にすることはできません",
	"ID number cannot be empty":                          "ID 番号を空にすることはできません",
	"Phone number can not be blank":                      "携帯電話番号を空にすることはできません",
	"The front of the ID card cannot be blank":           "ID の前面を空にすることはできません",
	"The back of the ID card cannot be blank":            "ID の裏側は空にできません",
	"Real name authentication already exists":            "実名認証は既に存在します",
	"Closed on weekends":                                 "週末は休業",
	"credential image must be required!":                 "支払い伝票の画像を空にすることはできません!",
	"The withdrawal amount format is incorrect":          "出金金額の形式が正しくありません",
	"The phone format is incorrect":                      "電話の形式が正しくありません",
	"The ID card format is incorrect":                    "ID 形式が正しくありません",
	"The ID number has been submitted for certification": "ID 番号が認証のために送信されました",
	"You have submitted real name authentication":        "あなたの実名認証情報は既に存在します",
	"Image upload failed":                                "画像のアップロードに失敗しました",
}
