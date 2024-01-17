import { encodeURIParam, decodeURIParam } from "@/utilities/paramsParsing";

describe('URI Encoding and Decoding', () => {
  it('should encode a simple string', () => {
    const input = 'Hello, World!';
    const encoded = encodeURIParam(input);
    expect(encoded).toBe('Hello,%20World!');
  });

  it('should encode a string with special characters', () => {
    const input = 'This is a string with spaces & special characters: / ? = #';
    const encoded = encodeURIParam(input);
    expect(encoded).toBe('This%20is%20a%20string%20with%20spaces%20&%20special%20characters:%20/%20?%20=%20#');
  });

  it('should decode an encoded string', () => {
    const encoded = 'This%20is%20an%20encoded%20string.';
    const decoded = decodeURIParam(encoded);
    expect(decoded).toBe('This is an encoded string.');
  });

  it('should handle non-encoded string in decoding', () => {
    const input = 'This is not encoded.';
    const decoded = decodeURIParam(input);
    expect(decoded).toBe('This is not encoded.');
  });
});
