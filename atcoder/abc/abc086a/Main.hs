
main = do
    -- map read . words でString -> [Int]の関数になっている
    [a,b] <- map read . words <$> getLine
    if (a * b) `mod` 2 == 0 then putStrLn "Even" else putStrLn "Odd"