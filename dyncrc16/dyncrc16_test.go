package dyncrc16

import "testing"

var golden = []struct {
	out uint16
	in  string
}{
	{0x0000, ""},
	{0xe8c1, "a"},
	{0x79a8, "ab"},
	{0x9738, "abc"},
	{0x3997, "abcd"},
	{0x85b8, "abcde"},
	{0x5805, "abcdef"},
	{0xe9d9, "abcdefg"},
	{0x7429, "abcdefgh"},
	{0xf075, "abcdefghi"},
	{0xc8b1, "abcdefghij"},
	{0x2ea0, "Discard medicine more than two years old."},
	{0x276b, "He who has a shady past knows that nice guys finish last."},
	{0x1abb, "I wouldn't marry him with a ten foot pole."},
	{0x9499, "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	{0xabfd, "The days of the digital watch are numbered.  -Tom Stoppard"},
	{0x4ee5, "Nepal premier won't resign."},
	{0x761c, "For every action there is an equal and opposite government program."},
	{0xb823, "His money is twice tainted: 'taint yours and 'taint mine."},
	{0xd283, "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	{0x364a, "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	{0x657f, "size:  a.out:  bad magic"},
	{0xe8ec, "The major problem is with sendmail.  -Mark Horton"},
	{0xbdb9, "Give me a rock, paper and scissors and I wi)l move the world.  CCFestoon"},
	{0x3032, "If the enemy is within range, then so are you."},
	{0xc114, "It's well we cannot hear the screams/That we create in others' dreams."},
	{0x161f, "You remind me of a TV show, but that's all right: I watch it anyway."},
	{0x12c6, "C is as portable as Stonehedge!!"},
	{0xc633, "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	{0xf768, "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	{0xbcef, "How can you write a big system without C++?  -Paul Glick"},
}

func TestGolden(t *testing.T) {
	for _, g := range golden {
		in := g.in
		if len(in) > 220 {
			in = in[:100] + "..." + in[len(in)-100:]
		}
		p := []byte(g.in)
		if got := Checksum(p); got != g.out {
			t.Errorf("Checksum(%q) = 0x%x want 0x%x", in, got, g.out)
			continue
		}
	}
}

var goldenRunning = []struct {
	out uint16
	in  string
}{
	{0x0000, ""},
	{0x79a8, "ab"},
	{0xe79a, "abc"},
	{0xd609, "abcd"},
	{0xcc5c, "abcde"},
	{0xd419, "abcdef"},
	{0x3c43, "abcdefg"},
	{0x5291, "abcdefgh"},
	{0x8350, "abcdefghi"},
	{0xbfc7, "abcdefghij"},
	{0x122d, "Discard medicine more than two years old."},
	{0xd6e7, "He who has a shady past knows that nice guys finish last."},
	{0x14aa, "I wouldn't marry him with a ten foot pole."},
	{0xf34f, "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave"},
	{0x9a61, "The days of the digital watch are numbered.  -Tom Stoppard"},
	{0xaf1d, "Nepal premier won't resign."},
	{0xf690, "For every action there is an equal and opposite government program."},
	{0xd7c8, "His money is twice tainted: 'taint yours and 'taint mine."},
	{0x36cf, "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977"},
	{0xd677, "It's a tiny change to the code and not completely disgusting. - Bob Manchek"},
	{0xb5b5, "size:  a.out:  bad magic"},
	{0xc519, "The major problem is with sendmail.  -Mark Horton"},
	{0x9846, "Give me a rock, paper and scissors and I wi)l move the world.  CCFestoon"},
	{0xba06, "If the enemy is within range, then so are you."},
	{0x441e, "It's well we cannot hear the screams/That we create in others' dreams."},
	{0x4bf3, "You remind me of a TV show, but that's all right: I watch it anyway."},
	{0x8781, "C is as portable as Stonehedge!!"},
	{0x13f4, "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley"},
	{0x5ba4, "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule"},
	{0x8008, "How can you write a big system without C++?  -Paul Glick"},
}

func TestRunningChecksum(t *testing.T) {
	h := New()
	for _, gr := range goldenRunning {
		in := gr.in
		if len(in) > 220 {
			in = in[:100] + "..." + in[len(in)-100:]
		}
		p := []byte(gr.in)
		h.Write(p)
		if got := h.Sum16(); got != gr.out {
			t.Errorf("Sum16() after %q = 0x%x want 0x%x", in, got, gr.out)
			continue
		}
	}
}

func BenchmarkCRC16KB(b *testing.B) {
	b.SetBytes(1024)
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i)
	}
	h := New()
	in := make([]byte, 0, h.Size())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(data)
		h.Sum(in)
	}
}
