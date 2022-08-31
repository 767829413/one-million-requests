package plan

func Handler1(payloads []interface{}) {
	for _, payload := range payloads {
		go func(payload interface{}) {
			// do something
			_ = payload
		}(payload)
	}
}
