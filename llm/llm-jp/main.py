import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

tokenizer = AutoTokenizer.from_pretrained("llm-jp/llm-jp-1.3b-v1.0")
model = AutoModelForCausalLM.from_pretrained(
    "llm-jp/llm-jp-1.3b-v1.0", device_map="auto", torch_dtype=torch.float16
)
text = "自然言語処理とは何か"
tokenized_input = tokenizer.encode(
    text, add_special_tokens=False, return_tensors="pt"
).to(model.device)
with torch.no_grad():
    output = model.generate(
        tokenized_input,
        max_new_tokens=20,
        do_sample=True,
        top_p=0.90,
        temperature=0.7,
    )[0]
print(tokenizer.decode(output))
