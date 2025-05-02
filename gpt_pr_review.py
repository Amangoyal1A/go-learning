import os
import subprocess
import openai

openai.api_key = os.getenv("OPENAI_API_KEY")

def get_git_diff():
    # Diff against the base branch (works in PRs)
    return subprocess.check_output(["git", "diff", "origin/main...HEAD"]).decode("utf-8")

def ask_gpt(diff):
    response = openai.ChatCompletion.create(
        model="gpt-4",
        messages=[
            {"role": "system", "content": "You're a senior Go engineer reviewing a pull request. Your job is to look for naming convention issues, bug-prone logic, and give helpful suggestions."},
            {"role": "user", "content": f"Review this Git diff:\n\n{diff}\n\nPlease point out any problems and suggest improvements."}
        ]
    )
    return response['choices'][0]['message']['content']

def main():
    diff = get_git_diff()
    if not diff.strip():
        print("No code changes to review.")
        return
    review = ask_gpt(diff)
    print("------ GPT REVIEW ------")
    print(review)

if __name__ == "__main__":
    main()
