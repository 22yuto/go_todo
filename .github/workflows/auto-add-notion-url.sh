name: Add Notion Ticket URL to PR

on:
  pull_request:
    types: [opened]

jobs:
  add-notion-url:
    runs-on: ubuntu-latest
    steps:
      - name: Check out code
        uses: actions/checkout@v3

      - name: Debug Notion Response
        run: |
        curl -X POST "https://api.notion.com/v1/pages" \
        -H "Authorization: Bearer ${{ secrets.NOTION_API_KEY }}" \
        -H "Content-Type: application/json" \
        -d '{
          "filter": {
            "property": "Status",
            "status": "In Progress"
          }
        }' | tee notion_debug_response.json
      - name: Fetch Notion Ticket URL
        id: fetch-notion
        run: |
          # Notion APIを利用してチケットのURLを取得
          # 必要に応じてカスタマイズ
          curl -X POST "https://api.notion.com/v1/pages" \
          -H "Authorization: Bearer ${{ secrets.NOTION_API_KEY }}" \
          -H "Content-Type: application/json" \
          -d '{
            "filter": {
              "property": "Status",
              "status": "In Progress"
            }
          }' > notion_ticket.json

          # URLを抽出
          export TICKET_URL=$(jq -r '.results[0].url' notion_ticket.json)
          echo "::set-output name=url::$TICKET_URL"

      - name: Update PR description
        env:
          PR_URL: ${{ steps.fetch-notion.outputs.url }}
        run: |
          gh pr edit ${{ github.event.pull_request.number }} \
          --body "$(echo -e "## チケット\n${PR_URL}\n\n$(cat README.md)")"
