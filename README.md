# Skillset Example

## Description

This code sample demonstrates building a Copilot Extension using the skillsets approach rather than a traditional agent. This extension is designed to generate random test and example data for a number of development purposes, by calling publicly available APIs.

### Architectural Model
- **Skillsets**: Define up to 5 API endpoints that Copilot can call directly. Copilot handles all AI interactions, prompt engineering, and response formatting.
- **Agents**: Provide full control over the interaction flow, including custom prompt crafting and specific LLM model selection.

![Architectural comparison between Skillsets and Agents](https://github.com/user-attachments/assets/9c5d6489-afb5-47c2-be73-2561d89dfde3)


### When to Choose Skillsets
Skillsets are ideal when you want to:
- Quickly integrate existing APIs or services without managing AI logic
- Focus purely on your service's core functionality
- Maintain consistent Copilot-style interactions without extensive development
- Get started with minimal infrastructure and setup

Use agents instead if you need:
- Complex custom interaction flows
- Specific LLM model control (using LLMs that aren't provided by the Copilot API)
- Custom prompt crafting
- Advanced state management

## Example Implementation

This extension showcases the skillset approach by providing three simple endpoints that generate random development data:
- Random commit messages
- URL shortener
- Random user data

## Getting Started
1. Clone the repository: 

```
git clone git@github.com:copilot-extensions/skillset-example.git
cd skillset-example
```

2. Install dependencies:

```
go mod tidy
```

## Usage

1. Start up ngrok with the port provided:

```
ngrok http http://localhost:8080
```

2. Set the environment variables (use the ngrok generated url for the `FDQN`)
3. Run the application:

```
go run .
```

## Accessing the Extension in Chat:

1. In the `Copilot` tab of your Application settings (`https://github.com/settings/apps/<app_name>/agent`)
- Set the app type to "Skillset"
- Specify the following skills
```
Name: random_commit_message
Inference description: Generates a random commit message
URL: https://<your ngrok domain>/random-commit-message
Parameters: { "type": "object" }
Return type: String
---
Name: shorten_url
Inference description: Shortens a long URL using the CleanURI service. Send a JSON body with the "url" field and receive the shortened URL in the response.
URL: https://<your ngrok domain>/shorten-url
Parameters: 
{
  "type": "object",
  "properties": {
    "url": {
      "type": "string",
      "description": "The long URL to be shortened."
    }
  },
  "required": ["url"]
}
Return type: String
---
Name: random_user
Inference description: Generates data for a random user
URL: https://<your ngrok domain>/random-user
Parameters: { "type": "object" }
Return type: String
```

2. In the `General` tab of your application settings (`https://github.com/settings/apps/<app_name>`)
- Set the `Callback URL` to anything (`https://github.com` works well for testing, in a real environment, this would be a URL you control)
- Set the `Homepage URL` to anything as above
3. Ensure your permissions are enabled in `Permissions & events` > 
- `Account Permissions` > `Copilot Chat` > `Access: Read Only`
4. Ensure you install your application at (`https://github.com/apps/<app_name>`)
5. Now if you go to `https://github.com/copilot` you can `@` your skillset extension using the name of your app.

## What can the bot do?

Here's some example things:

* `@skillset-example please create a random commit message`
* `@skillset-example please short this url https://github.com`
* `@skillset-example generate random user data`

## Implementation

This bot provides a passthrough to a couple of other APIs:

* For commit messages, https://whatthecommit.com/
* For shorting URLs, https://cleanuri.com/
* For user data, https://randomuser.me/

## Documentation
- [Using Copilot Extensions](https://docs.github.com/en/copilot/using-github-copilot/using-extensions-to-integrate-external-tools-with-copilot-chat)
- [About skillsets](https://docs.github.com/en/copilot/building-copilot-extensions/building-a-copilot-skillset-for-your-copilot-extension/about-copilot-skillsets)
- [About building Copilot Extensions](https://docs.github.com/en/copilot/building-copilot-extensions/about-building-copilot-extensions)
- [Set up process](https://docs.github.com/en/copilot/building-copilot-extensions/setting-up-copilot-extensions)
