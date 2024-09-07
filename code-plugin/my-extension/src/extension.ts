// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
import * as vscode from "vscode";
import { exec } from "child_process";
import { assert } from "console";

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed
export function activate(context: vscode.ExtensionContext) {
  // Use the console to output diagnostic information (console.log) and errors (console.error)
  // This line of code will only be executed once when your extension is activated
  console.log('Congratulations, your extension "my-extension" is now active!');

  // The command has been defined in the package.json file
  // Now provide the implementation of the command with registerCommand
  // The commandId parameter must match the command field in package.json
  const disposable = vscode.commands.registerCommand(
    "my-extension.helloWorld",
    () => {
      // The code you place here will be executed every time your command is executed
      // Display a message box to the user
      vscode.window.showInformationMessage("Hello World from my-extension!");
    },
  );
  const myDisposable = vscode.commands.registerCommand(
    "my-extension.asumi",
    () => {
      exec("az account show", (_err, stdout, _stderr) => {
        console.log(stdout);
        vscode.window.showInformationMessage(stdout);
      });
    },
  );
  const hoge = vscode.commands.registerCommand(
    "my-extension.hoge",
    async () => {
      // const item = await vscode.window.showQuickPick(
      //   ["hoge", "fuga", "asumi"],
      //   {
      //     placeHolder: "select items",
      //     canPickMany: false,
      //   },
      // );
      let ranges: Array<string> = [];
      const num = await vscode.window
        .showInputBox({ title: "input number" })
        .then((val) => {
          return val;
        });

      for (let i = 0; i < Number(num); ++i) {
        ranges.push(i.toString());
      }

      await vscode.window
        .showQuickPick(ranges, {
          placeHolder: "select items",
          canPickMany: false,
        })
        .then((selected) => {
          vscode.window.showInformationMessage(`selected: ${selected}`);
        });
      vscode.window
        .showInformationMessage("is Ok?", { modal: true }, { title: "Yes" })
        .then((answer) => {
          console.log(answer?.title);
          if (answer?.title === "Yes") {
            console.error("Yes");
          } else if (answer?.title === "No") {
            console.error("No");
          } else {
            console.log(answer?.title);
          }
        });
    },
  );

  context.subscriptions.push(disposable);
  context.subscriptions.push(myDisposable);
  context.subscriptions.push(hoge);
}

// This method is called when your extension is deactivated
export function deactivate() {}
