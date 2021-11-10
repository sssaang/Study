package com.sssaang.springboot.tutorial.tutorial.datasource.mock

import org.assertj.core.api.Assertions
import org.junit.jupiter.api.Test

internal class MockBankDataSourceTest {
    private val mockDataSource = MockBankDataSource()
    @Test
    fun `should provide a collection of banks`() {
        val banks = mockDataSource.retrieveBanks()

        Assertions.assertThat(banks).isNotEmpty
        Assertions.assertThat(banks.size).isGreaterThan(0)
    }

    @Test
    fun `should provide a some mock data`() {
        val banks = mockDataSource.retrieveBanks()

        Assertions.assertThat(banks).allMatch { it.accountNumber.isNotBlank() }
        Assertions.assertThat(banks).anyMatch { it.trust != 0.0 }
        Assertions.assertThat(banks).anyMatch { it.transactionFee != 0 }
    }
}