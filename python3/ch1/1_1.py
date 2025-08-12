def is_unique(string: str) -> bool:
    # can be the input string be empty or null?

    # traverse the string
    # save every key in a hashmap
    # if a colition is found return False
    # return True when the traverse is complete

    # error 1: didn't handle the last letter
    # solution 1: hashmap value should be set as a `truthy` value, not 0

    if not string:
        return True

    hashmap: dict[str, int] = {}
    for letter in string:
        if hashmap.get(letter):
            return False
        hashmap[letter] = 1
    return True



if __name__ == "__main__":
    import unittest

    tests = [("asdfg", True), ("", True), ("asdfa", False), (" ", True), ("  ", False)]

    class Test(unittest.TestCase):
        def test_loop(self):
            for t in tests:
                result = is_unique(t[0])
                self.assertEqual(
                        first=result,
                        second=t[1],
                        msg=f"is_unique(\"{t[0]}\") is not {t[1]}. got={result}",
                        )

    unittest.main()
