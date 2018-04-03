# slack_image_bot
 This bot will reply the image URL to the mentions

<div style="text-align:center"><img src="https://qiita-image-store.s3.amazonaws.com/0/81341/9580fe00-7f06-e7cc-d268-5706f033b750.gif" height='300' width='500'></div>
 
## Set up
### 1. Google custom search
- Get google custom search API Key
<br>[Google Custom search](https://developers.google.com/custom-search/json-api/v1/overview?hl=en)
- Plese add search engine and get engine key
<br>[Add search engine](https://cse.google.com/manage/all)
### 2. Slack bot 
- Create slack bot and get slack token
<br>[Create slack bot](https://slack.com/customize/slackbot)
### 3. Add environment value
- SLACK_API_TOKEN
- CUSTOM_SEARCH_KEY
- CUSTOM_SEARCH_ENGINE_ID
## Try!!
`go run main.go` or `make run`
