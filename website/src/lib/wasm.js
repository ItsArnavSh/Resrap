import './wasm_exec.js';

let wasmReady = false;
let goInstance = null;

export async function loadWasm() {
	if (wasmReady) {
		console.log('WASM already loaded');
		return;
	}

	console.log('🚀 Starting WASM load...');

	try {
		console.log('📋 Checking Go availability:', typeof Go);
		if (typeof Go === 'undefined') {
			throw new Error('Go constructor not found');
		}

		console.log('🔧 Creating Go instance...');
		goInstance = new Go();
		console.log('✅ Go instance created:', goInstance);
		console.log('📦 Import object keys:', Object.keys(goInstance.importObject));

		console.log('📡 Fetching main.wasm...');
		const wasmResponse = await fetch('/main.wasm');
		console.log('📊 Response status:', wasmResponse.status, wasmResponse.statusText);

		if (!wasmResponse.ok) {
			throw new Error(`HTTP ${wasmResponse.status}: ${wasmResponse.statusText}`);
		}

		console.log('🔄 Converting to ArrayBuffer...');
		const wasmArrayBuffer = await wasmResponse.arrayBuffer();
		console.log('📏 WASM size:', wasmArrayBuffer.byteLength, 'bytes');

		console.log('⚡ Instantiating WASM module...');
		const wasmModule = await WebAssembly.instantiate(wasmArrayBuffer, goInstance.importObject);
		console.log('✅ WASM instantiated:', wasmModule);

		console.log('🏃 Running Go program...');
		goInstance.run(wasmModule.instance);

		wasmReady = true;
		console.log('🎉 WASM fully loaded and running!');

		// Check if functions are available
		setTimeout(() => {
			console.log('🔍 Checking global functions:');
			console.log('- goTest:', typeof globalThis.goTest);
			console.log(
				'- Available keys on globalThis containing "go":',
				Object.keys(globalThis).filter((key) => key.toLowerCase().includes('go'))
			);
		}, 100);
	} catch (error) {
		console.error('❌ WASM loading failed at step:', error.message);
		console.error('🔍 Full error:', error);
		console.error('📚 Error stack:', error.stack);
		throw error;
	}
}

export function unloadWasm() {
	if (goInstance && goInstance.exit) {
		goInstance.exit();
	}
	goInstance = null;
	wasmReady = false;
}
