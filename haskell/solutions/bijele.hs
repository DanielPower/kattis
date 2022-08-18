bijele :: (Int, Int, Int, Int, Int, Int) -> [Int]
bijele (kings, queens, rooks, bishops, knights, pawns) = [1 - kings, 1 - queens, 2 - rooks, 2 - bishops, 2 - knights, 8 - pawns]

readInput :: String -> [Int]
readInput = map read . words

writeOutput :: [Int] -> String
writeOutput = unwords . map show

solve :: [Int] -> [Int]
solve values = bijele (head values, values !! 1, values !! 2, values !! 3, values !! 4, values !! 5)

main = interact (writeOutput . solve . readInput)
