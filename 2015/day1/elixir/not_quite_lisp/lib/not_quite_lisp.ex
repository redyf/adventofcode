defmodule NotQuiteLisp do
  def read_file_and_loop(filepath) do
    case File.read(filepath) do
      {:ok, contents} ->
        count =
          contents
          |> String.split("", trim: true)
          |> Enum.reduce(0, fn char, acc ->
            acc + count_parentheses(char)
          end)

        {:ok, count}

      {:error, reason} ->
        {:error, "Error reading the file: #{reason}"}
    end
  end

  defp count_parentheses(char) do
    case char do
      "(" -> 1
      ")" -> -1
      _ -> 0
    end
  end
end
