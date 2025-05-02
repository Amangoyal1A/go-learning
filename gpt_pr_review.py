import os
import subprocess
from openai import OpenAI

# Auth + client setup
client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))
base_branch = os.getenv("BASE_BRANCH", "main")

def get_git_diff():
    subprocess.run(["git", "fetch", "origin", base_branch], check=True)
    return subprocess.check_output(["git", "diff", f"origin/{base_branch}...HEAD"]).decode("utf-8")

def ask_gpt(diff):
    response = client.chat.completions.create(
        model="gpt-4",
        messages=[
            {"role": "system", "content": "You're a senior Go engineer reviewing a pull request. Point out naming convention issues, potential bugs, and optimization ideas."},
            {"role": "user", "content": f"Review this Git diff:\n\n{diff}"}
        ]
    )
    return response.choices[0].message.content

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
