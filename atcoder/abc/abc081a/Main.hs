
main = do
    input <- getLine
    print . length . filter (=='1') $ input