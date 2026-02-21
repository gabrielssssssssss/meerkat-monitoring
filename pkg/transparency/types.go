package transparency

type Tree struct {
	TreeSize          int64  `json:"tree_size"`
	Timestamp         int64  `json:"timestamp"`
	Sha256RootHash    string `json:"sha256_root_hash"`
	TreeHeadSignature string `json:"tree_head_signature"`
}

type Entries struct {
	Entries []struct {
		LeafInput string `json:"leaf_input"`
		ExtraData string `json:"extra_data"`
	} `json:"entries"`
}
