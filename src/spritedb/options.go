package spritedb

type OptFunc func(opts *Options)

type Options struct {
	Encoder DataEncoder
	Decoder DataDecoder
	DBName  string
}

func WithEncoder(encoder DataEncoder) OptFunc {
	return func(opts *Options) {
		opts.Encoder = encoder
	}
}

func WithDecoder(decoder DataDecoder) OptFunc {
	return func(opts *Options) {
		opts.Decoder = decoder
	}
}
