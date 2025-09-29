declare global {
  var generateText: ((grammar: string, startpoint: string, seed: number, tokenlen: number) => string) | undefined;
  var Go: new () => any;
  var goTest: () => string;
}

export {};