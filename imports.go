// Copyright 2021 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package main

import (
	"fmt"
	"math"

	"github.com/ChainSafe/gossamer/lib/runtime"
	"github.com/ChainSafe/gossamer/pkg/scale"
	"github.com/wasmerio/wasmer-go/wasmer"
)

const (
	validateSignatureFail = "failed to validate signature"
)

//export ext_logging_log_version_1
func ext_logging_log_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{}, nil
}

//export ext_logging_max_level_version_1
func ext_logging_max_level_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_transaction_index_index_version_1
func ext_transaction_index_index_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{}, nil
}

//export ext_transaction_index_renew_version_1
func ext_transaction_index_renew_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_sandbox_instance_teardown_version_1
func ext_sandbox_instance_teardown_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_sandbox_instantiate_version_1
func ext_sandbox_instantiate_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_sandbox_invoke_version_1
func ext_sandbox_invoke_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_sandbox_memory_get_version_1
func ext_sandbox_memory_get_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_sandbox_memory_new_version_1
func ext_sandbox_memory_new_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_sandbox_memory_set_version_1
func ext_sandbox_memory_set_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_sandbox_memory_teardown_version_1
func ext_sandbox_memory_teardown_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_ed25519_generate_version_1
func ext_crypto_ed25519_generate_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_ed25519_public_keys_version_1
func ext_crypto_ed25519_public_keys_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_ed25519_sign_version_1
func ext_crypto_ed25519_sign_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_ed25519_verify_version_1
func ext_crypto_ed25519_verify_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_secp256k1_ecdsa_recover_version_1
func ext_crypto_secp256k1_ecdsa_recover_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_secp256k1_ecdsa_recover_version_2
func ext_crypto_secp256k1_ecdsa_recover_version_2(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return ext_crypto_secp256k1_ecdsa_recover_version_1(env, args)
}

//export ext_crypto_ecdsa_verify_version_2
func ext_crypto_ecdsa_verify_version_2(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_secp256k1_ecdsa_recover_compressed_version_1
func ext_crypto_secp256k1_ecdsa_recover_compressed_version_1(env interface{},
	args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_secp256k1_ecdsa_recover_compressed_version_2
func ext_crypto_secp256k1_ecdsa_recover_compressed_version_2(env interface{},
	args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_sr25519_generate_version_1
func ext_crypto_sr25519_generate_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_sr25519_public_keys_version_1
func ext_crypto_sr25519_public_keys_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_sr25519_sign_version_1
func ext_crypto_sr25519_sign_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_sr25519_verify_version_1
func ext_crypto_sr25519_verify_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_sr25519_verify_version_2
func ext_crypto_sr25519_verify_version_2(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_start_batch_verify_version_1
func ext_crypto_start_batch_verify_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_crypto_finish_batch_verify_version_1
func ext_crypto_finish_batch_verify_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_trie_blake2_256_root_version_1
func ext_trie_blake2_256_root_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_trie_blake2_256_ordered_root_version_1
func ext_trie_blake2_256_ordered_root_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_trie_blake2_256_ordered_root_version_2
func ext_trie_blake2_256_ordered_root_version_2(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_trie_blake2_256_verify_proof_version_1
func ext_trie_blake2_256_verify_proof_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_misc_print_hex_version_1
func ext_misc_print_hex_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_misc_print_num_version_1
func ext_misc_print_num_version_1(_ interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_misc_print_utf8_version_1
func ext_misc_print_utf8_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_misc_runtime_version_version_1
func ext_misc_runtime_version_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	version := runtime.Version{
		SpecName:         []byte("westend"),
		ImplName:         []byte("gossamer"),
		AuthoringVersion: 789,
		SpecVersion:      1,
		ImplVersion:      1,
		APIItems: []runtime.APIItem{
			{
				Name: [8]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
				Ver:  0,
			},
			{
				Name: [8]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
				Ver:  0,
			},
			{
				Name: [8]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
				Ver:  0,
			},
			{
				Name: [8]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8},
				Ver:  0,
			},
		},
		TransactionVersion: 1,
		StateVersion:       1,
	}

	// Note the encoding contains all the latest Core_version fields as defined in
	// https://spec.polkadot.network/#defn-rt-core-version
	// In other words, decoding older version data with missing fields
	// and then encoding it will result in a longer encoding due to the
	// extra version fields. This however remains compatible since the
	// version fields are still encoded in the same order and an older
	// decoder would succeed with the longer encoding.
	encodedData, err := scale.Marshal(version)
	if err != nil {
		return []wasmer.Value{wasmer.NewI64(0)}, nil
	}

	out, err := toWasmMemoryOptional(env.(*Runtime), encodedData)
	if err != nil {
		return []wasmer.Value{wasmer.NewI64(0)}, nil
	}

	return []wasmer.Value{wasmer.NewI64(out)}, nil
}

// toWasmMemoryOptional scale encodes the byte slice `data`, writes it to wasm memory
// and returns the corresponding 64 bit pointer size.
func toWasmMemoryOptional(context *Runtime, data []byte) (
	pointerSize int64, err error) {
	var optionalSlice *[]byte
	if data != nil {
		optionalSlice = &data
	}

	encoded, err := scale.Marshal(optionalSlice)
	if err != nil {
		return 0, err
	}

	return toWasmMemory(context, encoded)
}

// toWasmMemory copies a Go byte slice to wasm memory and returns the corresponding
// 64 bit pointer size.
func toWasmMemory(context *Runtime, data []byte) (
	pointerSize int64, err error) {
	size := uint32(len(data))

	ptr, err := context.allocator.Allocate(size)
	if err != nil {
		return 0, fmt.Errorf("allocating: %w", err)
	}

	memory := context.mem.Data()

	if uint32(len(memory)) < ptr+size {
		panic(fmt.Sprintf("length of memory is less than expected, want %d have %d", ptr+size, len(memory)))
	}

	copy(memory[ptr:ptr+size], data)
	pointerSize = toPointerSize(ptr, size)
	return pointerSize, nil
}

// toPointerSize converts an uint32 pointer and uint32 size
// to an int64 pointer size.
func toPointerSize(ptr, size uint32) (pointerSize int64) {
	return int64(ptr) | (int64(size) << 32)
}

//export ext_default_child_storage_read_version_1
func ext_default_child_storage_read_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_clear_version_1
func ext_default_child_storage_clear_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_clear_prefix_version_1
func ext_default_child_storage_clear_prefix_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_exists_version_1
func ext_default_child_storage_exists_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_get_version_1
func ext_default_child_storage_get_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_next_key_version_1
func ext_default_child_storage_next_key_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_root_version_1
func ext_default_child_storage_root_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_set_version_1
func ext_default_child_storage_set_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_storage_kill_version_1
func ext_default_child_storage_storage_kill_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_storage_kill_version_2
func ext_default_child_storage_storage_kill_version_2(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_default_child_storage_storage_kill_version_3
func ext_default_child_storage_storage_kill_version_3(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_allocator_free_version_1
func ext_allocator_free_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	runtimeCtx := env.(*Runtime)
	addr, ok := args[0].Unwrap().(int32)
	if !ok {
		fmt.Printf("[ext_allocator_free_version_1] error addr cannot be converted to int32\n")
	}

	// Deallocate memory
	err := runtimeCtx.allocator.Deallocate(uint32(addr))
	if err != nil {
		panic(err)
	}
	return nil, nil
}

//export ext_allocator_malloc_version_1
func ext_allocator_malloc_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	size, ok := args[0].Unwrap().(int32)
	if !ok {
		fmt.Printf("[ext_allocator_malloc_version_1] error addr cannot be converted to int32\n")
	}

	ctx := env.(*Runtime)

	// Allocate memory
	res, err := ctx.allocator.Allocate(uint32(size))
	if err != nil {
		panic(err)
	}

	castedRes, err := safeCastInt32(res)
	if err != nil {
		//logger.Errorf("failed to safely cast pointer: %s", err)
		return []wasmer.Value{wasmer.NewI32(0)}, err
	}

	return []wasmer.Value{wasmer.NewI32(castedRes)}, nil

	//return []wasmer.Value{wasmer.NewI32(int32(res))}, nil
}

func safeCastInt32(value uint32) (int32, error) {
	if value > math.MaxInt32 {
		return 0, fmt.Errorf("out of bounds")
	}
	return int32(value), nil
}

//export ext_hashing_blake2_128_version_1
func ext_hashing_blake2_128_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_hashing_blake2_256_version_1
func ext_hashing_blake2_256_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_hashing_keccak_256_version_1
func ext_hashing_keccak_256_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_hashing_sha2_256_version_1
func ext_hashing_sha2_256_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_hashing_twox_256_version_1
func ext_hashing_twox_256_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_hashing_twox_128_version_1
func ext_hashing_twox_128_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_hashing_twox_64_version_1
func ext_hashing_twox_64_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_index_set_version_1
func ext_offchain_index_set_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_local_storage_clear_version_1
func ext_offchain_local_storage_clear_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_is_validator_version_1
func ext_offchain_is_validator_version_1(env interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_local_storage_compare_and_set_version_1
func ext_offchain_local_storage_compare_and_set_version_1(env interface{},
	args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_local_storage_get_version_1
func ext_offchain_local_storage_get_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_local_storage_set_version_1
func ext_offchain_local_storage_set_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_network_state_version_1
func ext_offchain_network_state_version_1(env interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_random_seed_version_1
func ext_offchain_random_seed_version_1(env interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_submit_transaction_version_1
func ext_offchain_submit_transaction_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_timestamp_version_1
func ext_offchain_timestamp_version_1(_ interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_sleep_until_version_1
func ext_offchain_sleep_until_version_1(_ interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_http_request_start_version_1
func ext_offchain_http_request_start_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_offchain_http_request_add_header_version_1
func ext_offchain_http_request_add_header_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_append_version_1
func ext_storage_append_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_changes_root_version_1
func ext_storage_changes_root_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_clear_version_1
func ext_storage_clear_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_clear_prefix_version_1
func ext_storage_clear_prefix_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_clear_prefix_version_2
func ext_storage_clear_prefix_version_2(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_exists_version_1
func ext_storage_exists_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_get_version_1
func ext_storage_get_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_next_key_version_1
func ext_storage_next_key_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_read_version_1
func ext_storage_read_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_root_version_1
func ext_storage_root_version_1(env interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_root_version_2
func ext_storage_root_version_2(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_set_version_1
func ext_storage_set_version_1(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_start_transaction_version_1
func ext_storage_start_transaction_version_1(env interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_rollback_transaction_version_1
func ext_storage_rollback_transaction_version_1(env interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}

//export ext_storage_commit_transaction_version_1
func ext_storage_commit_transaction_version_1(env interface{}, _ []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(0)}, nil
}
