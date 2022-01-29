defmodule WordCount do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence
    |> String.split(~r{[[:blank:]_]})
    |> Enum.reduce(%{}, fn(sentence_part, acc) ->
      matches = Regex.run(~r{[[:alnum:]\-]+}u, sentence_part)

      case matches do
        nil -> acc
        _ -> count_word(acc, matches)
      end
    end)
  end

  @spec count_word(map, [binary()]) :: map
  defp count_word(acc, matches) do
    word = matches |> List.first() |> String.downcase()
    Map.update(acc, word, 1, fn value -> value + 1 end)
  end
end
