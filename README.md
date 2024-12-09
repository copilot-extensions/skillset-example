> [!NOTE]
> Copilot Extensions are in public preview and may be subject to change. Pre-release terms apply. You must have a GitHub Copilot license of any type to use Copilot Extensions.

# Skillset Example
Learn to build a Copilot Extension as a skillset. Just define your API endpoints and let Copilot handle the complex AI logic. This sample demonstrates the lighter, faster alternative to building a traditional agent.

### Architecture
- **Skillsets**: Define up to 5 API endpoints that Copilot can call directly. Copilot handles all AI interactions, prompt engineering, and response formatting.
- **Agents**: Provide full control over the interaction flow, including custom prompt crafting and specific LLM model selection.

![Architectural comparison between Skillsets and Agents](https://github.com/user-attachments/assets/9c5d6489-afb5-47c2-be73-2561d89dfde3)

### Key Differences
|  | Skillsets | Agents |
|--------|-----------|--------|
| **Development Time** | Minutes to hours | Several days to weeks |
| **Implementation** | Configure API endpoints only. Quicker if you have existing APIs. | Custom AI logic, prompt engineering, interaction handling. |
| **Best For** | • Quick API integrations<br>• Basic service functionality<br>• Standard Copilot interactions<br>• Minimal infrastructure needs | • Complex interaction flows<br>• Custom LLM model control<br>• Custom prompt crafting<br>• Advanced state management |
| **Control Level** | Copilot platform handles AI interactions | Full control over entire workflow |
| **Data Volume** | Good for any volume, if the API handles efficient search/retrieval | Good for any volume, but requires robust search/retrieval at larger sizes |

## Example Prompts

This extension turns three public APIs into functions that can be used in Copilot Chat":
```typescript
@extension generate a commit message
→ "Whatever will be, will be 8{)"

@extension generate 2 short lorem ipsum paragraphs
→ [HTML-formatted Lorem Ipsum text]

@extension generate random user data
→ {
  "name": "Miss Mona",
  "email": "mona.lisa@github.com",
  ...
}
```

## Quick Start
### 1. Set Up Repository
```
git clone git@github.com:copilot-extensions/skillset-example.git
cd skillset-example
go mod tidy
```

### 2. Make It Accessible
1. Run your app:
```
go run .
```

2. Then choose one method to expose it and copy the URL for your GitHub App config:
* **VS Code**: Set port 8080 public, copy Local Address
* **Codespaces**: Set port 8080 public, copy URL
* **Local with ngrok**:
  1. Install ngrok from ngrok.com
  2. Start tunnel: `ngrok http 8080`
  3. Copy the Forwarding URL
 
> Note: This step is for development and testing only. For production use, deploy your extension to a public server with proper infrastructure and security.

### 3. Configure GitHub App
1. Create a new GitHub App at `github.com/settings/apps/new` or go to an existing one
2. In the Copilot tab:
   - Set type to "Skillset"
   - Define the three skills below, using your URL from step 2:
```
Name: random_commit_message
Inference description: Generates a random commit message
URL: https://<your ngrok domain>/random-commit-message
Parameters: { "type": "object" }
Return type: String
---
Name: random_lorem_ipsum 
Inference description: Generates a random Lorem Ipsum text.  Responses should have html tags present.
URL: https://<your ngrok domain>/random-lorem-ipsum
Parameters: 
{
   "type": "object",
   "properties": {
      "number_of_paragraphs": {
         "type": "number",
         "description": "The number of paragraphs to be generated.  Must be between 1 and 10 inclusive"
      },
      "paragraph_length": {
         "type": "string",
         "description": "The length of each paragraph.  Must be one of \"short\", \"medium\", \"long\", or \"verylong\""
      }
   }
}
Return type: String
---
Name: random_user
Inference description: Generates data for a random user
URL: https://<your ngrok domain>/random-user
Parameters: { "type": "object" }
Return type: String
```
3. In the `General` tab of your GitHub App settings (`https://github.com/settings/apps/<app_name>`)
   - **Homepage URL**: You can set to any URL for quick testing (like `https://github.com`).
   - **Callback URL**: Only needed if your extension uses OAuth for authentication. Since this example doesn't use OAuth, set to any URL (like `https://github.com`)
4. Ensure your permissions are enabled in `Permissions & events` > `Account Permissions` > `Copilot Chat` > `Access: Read Only`
5. Install your app at `https://github.com/apps/<app_name>`

### 4. Try It Out
1. Open Copilot Chat on github.com or in any supported IDE
3. Send a test prompt like `@test-data-bot generate a commit message`
> Note: It may take a few seconds for your extension to appear. Try refreshing if needed.

## Documentation
- [Using Copilot Extensions](https://docs.github.com/en/copilot/using-github-copilot/using-extensions-to-integrate-external-tools-with-copilot-chat)
- [About skillsets](https://docs.github.com/en/copilot/building-copilot-extensions/building-a-copilot-skillset-for-your-copilot-extension/about-copilot-skillsets)
- [About building Copilot Extensions](https://docs.github.com/en/copilot/building-copilot-extensions/about-building-copilot-extensions)
- [Set up process](https://docs.github.com/en/copilot/building-copilot-extensions/setting-up-copilot-extensions)
