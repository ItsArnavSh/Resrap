declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
	
	var generateText: ((grammar: string, startpoint: string, seed: number, tokenlen: number) => string) | undefined;
	var Go: new () => any;
	var goTest: () => string;
}

export {};
