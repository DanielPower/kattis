tarifa :: Int -> Int -> [Int] -> Int
tarifa x n p = x * (n + 1) - sum p

readInput :: String -> [Int]
readInput = map read . words

solve :: [Int] -> Int
solve values = tarifa (head values) (head (tail values)) (tail (tail values))

main = interact (show . solve . readInput)
