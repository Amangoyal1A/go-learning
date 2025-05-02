import os
import subprocess
import openai

openai.api_key = os.getenv("OPENAI_API_KEY")
base_branch = os.getenv("BASE_BRANCH", "main")

def get_git_diff():
    # Make sure we have the latest version of the base branch
    subprocess.run(["git", "fetch", "origin", base_branch], check=True)
    return subprocess.check_output(["git", "diff", f"origin/{base_branch}...HEAD"]).decode("utf-8")

def ask_gpt(diff):
    response = openai.ChatCompletion.create(
        model="gpt-4",
        messages=[
            {"role": "system", "content": "You're a senior Go engineer reviewing a pull request. Review the code diff and suggest naming improvements, detect bugs, and recommend optimizations."},
            {"role": "user", "content": f"Review this Git diff:\n\n{diff}\n\nCall out issues, bugs, naming problems, and suggest improvements."}
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
