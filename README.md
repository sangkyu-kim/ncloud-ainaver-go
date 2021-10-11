## Go SDK AI NAVER API Client for NAVER CLOUD PLATFROM

### Examples

```go
client := NewClient(os.Getenv("AINAVER_CLIENT_ID"), os.Getenv("AINAVER_SERCRET_KEY"))
ctx := context.Background()

response, err := client.PostTranslation(ctx, &PostTranslation{
    Source:    "en",
    Target:    "ko",
    Text:      "glad to meet you",
    Honorific: true,
})

if err != nil {
    log.Fatal(err)
}

log.Printf("Output: %+v", response)

if response.Message.Result.TranslatedText != "만나서 반가워요." {
    t.Fail()
}
```