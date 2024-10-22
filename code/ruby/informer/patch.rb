module Informers
  class PreTrainedTokenizer
    def self.test
      puts "test"
    end

    def self.from_pretrained(
      pretrained_model_name_or_path,
      quantized: true,
      progress_callback: nil,
      config: nil,
      cache_dir: nil,
      local_files_only: false,
      revision: "main",
      legacy: nil,
      **kwargs
    )
      tokenizer_json, tokenizer_config = load_tokenizer(
        pretrained_model_name_or_path,
        quantized:,
        progress_callback:,
        config:,
        cache_dir:,
        local_files_only:,
        revision:,
        legacy:
      )

      # Some tokenizers are saved with the "Fast" suffix, so we remove that if present.
      config = tokenizer_config["tokenizer_class"]
      puts config
      tokenizer_name = tokenizer_config["tokenizer_class"]&.delete_suffix("Fast") || "PreTrainedTokenizer"

      puts "TOKENIZER_CLASS_MAPPING"

      cls = TOKENIZER_CLASS_MAPPING[tokenizer_name]
      if !cls
        warn "Unknown tokenizer class #{tokenizer_name.inspect}, attempting to construct from base class."
        cls = PreTrainedTokenizer
      end

      cls.new(tokenizer_json, tokenizer_config)
    end
  end
end
