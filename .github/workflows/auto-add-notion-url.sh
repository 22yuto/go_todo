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

      - name: Install jq and gh CLI
        run: |
          sudo apt-get update && sudo apt-get install -y jq gh

      - name: Extract Ticket ID from PR title
        id: extract-ticket
        run: |
          # PRタイトルからTICKET_IDを抽出
          TICKET_ID=$(echo "${{ github.event.pull_request.title }}" | grep -oE '[A-Z]+-[0-9]+')
          echo "TICKET_ID=$TICKET_ID" >> $GITHUB_ENV

      - name: Fetch Notion Ticket URL
        id: fetch-notion
        env:
          NOTION_API_KEY: ${{ secrets.NOTION_API_KEY }}
          DATABASE_ID: ${{ secrets.NOTION_DATABASE_ID }}
          TICKET_ID: ${{ env.TICKET_ID }}
        run: |
          # TICKET_IDでデータベースをクエリ
          curl -X POST "https://api.notion.com/v1/databases/$DATABASE_ID/query" \
          -H "Authorization: Bearer $NOTION_API_KEY" \
          -H "Content-Type: application/json" \
          -d "{
            \"filter\": {
              \"property\": \"Ticket ID\",
              \"text\": {
                \"equals\": \"$TICKET_ID\"
              }
            }
          }" > notion_ticket.json

          # 結果からURLを抽出
          export TICKET_URL=$(jq -r '.results[0].url' notion_ticket.json)
          echo "TICKET_URL=$TICKET_URL" >> $GITHUB_ENV

      - name: Update PR description
        env:
          TICKET_URL: ${{ env.TICKET_URL }}
        run: |
          if [ -n "$TICKET_URL" ]; then
            gh pr edit ${{ github.event.pull_request.number }} \
            --body "$(echo -e "## チケット\n${TICKET_URL}\n\n$(cat README.md)")"
          else
            echo "No matching ticket found in Notion."
          fi
