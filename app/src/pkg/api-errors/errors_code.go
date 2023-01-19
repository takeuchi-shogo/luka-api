package apierrors

var (
	// 400
	BadRequest       = newBadRequest("bad_request", "リクエストを正常に処理できませんでした")
	Invalid          = newBadRequest("invalid", "このリクエストは無効です")
	InvalidParameter = newBadRequest("invalid_parameter", "無効なパラメータです")
	InvalidQuery     = newBadRequest("invalid_query", "無効なクエリです")
	NotDownload      = newBadRequest("not_download", "ダウンロードに失敗しました")
	NotUpload        = newBadRequest("not_upload", "アップロードに失敗しました")
	ParseError       = newBadRequest("parse_error", "サーバーがリクエスト本文を解析できません")
	Required         = newBadRequest("required", "リクエストに必要な情報がありません")
	UnknownApi       = newBadRequest("unknown_api", "リクエストが呼び出している API が認識されていません")

	// 401
	Unauthorized = newUnauthorized("unauthorized", "ユーザーは要求を行う権限がありません")
	// AuthError    = newAuthError("auth_error", "")

	// 402
	// PaymentRequired = newPaymentRequired("payment_required", "")

	// 403
	// Forbidden = newForbidden("forbidden", "")

	// 404
	NotFound = newNotFound("not_found", "要求に関連付けられたリソースが見つからなかったため、要求された操作は失敗しました")

	// 429
	TooManyRequests = newTooManyRequests("too_many_request", "特定の期間内に送信されたリクエストが多すぎます")

	// 500
	InternalServerError = newInternalServerError("internal_server_error", "内部エラーのため、要求は失敗しました")

	// 503
	ServiceUnavailable = newServiceUnavailable("service_unavailable", "バックエンド エラーが発生しました")
)
