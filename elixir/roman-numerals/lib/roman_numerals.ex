defmodule RomanNumerals do
  @main_numbers %{
    1 => "I",
    5 => "V",
    10 => "X",
    50 => "L",
    100 => "C",
    500 => "D",
    1000 => "M"
  }

  @doc """
  Convert the number to a roman number.
  """
  @spec numeral(pos_integer) :: String.t()
  def numeral(number) do
    convert_to_roman(number)
  end

  def convert_to_roman(number) do
    cond do
      number > 1 && number <= 3 ->
        String.duplicate(Map.get(@main_numbers, 1), number)

      number == 4 ->
        Map.get(@main_numbers, 1) <> Map.get(@main_numbers, 5)

      number > 5 && number < 9 ->
        tmp = number - 5
        part = String.duplicate(Map.get(@main_numbers, 1), tmp)
        Map.get(@main_numbers, 5) <> part

      number == 9 ->
        Map.get(@main_numbers, 1) <> Map.get(@main_numbers, 10)

      number > 10 && number < 40 ->
        [tens, rest] = to_string(number) |> String.split("", trim: true)
        String.duplicate(Map.get(@main_numbers, 10), String.to_integer(tens))
          <> convert_to_roman(String.to_integer(rest))

      true ->
        Map.get(@main_numbers, number)
    end
  end
end
