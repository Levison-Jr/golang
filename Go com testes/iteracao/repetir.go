package iteracao

// Repetir recebe um caractere e um inteiro n, e retorna uma string
// correspondente à repetição do caractere n vezes.
func Repetir(caractere string, quantidadeDeRepeticoes int) string {
	var repeticoes string
	for i := 0; i < quantidadeDeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
