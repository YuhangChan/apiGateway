namespace go kitex.service

struct Request{
    1: i32 operand_1 (api.body="operand_1")
    2: i32 operand_2 (api.body="operand_2")
}

struct Response{
    1: bool success (api.body="success")
    2: string message (api.body="message")
    3: i32 data (api.body="data")
}

service FirstLevelCalService{
    Response Add(1: Request request)(api.post="/gateway/FirstLevelCalService/add")
    Response Sub(1: Request request)(api.post="/gateway/FirstLevelCalService/sub")
}

