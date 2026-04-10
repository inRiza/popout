import { StyleSheet, View, Text, TextInput } from 'react-native';
import Feather from '@expo/vector-icons/Feather';

export default function Input({
    label,
    labelColor = 'black',
    placeholder = 'Enter your..',
    styleView,
    iconsName,
}: { label?: string, labelColor: string, placeholder: string, styleView?: any, iconsName?: any}) {

    return (
        <View style={[styles.container, styleView]}>
            <Text style={[styles.label, {
                color: labelColor,
            }]}>{label}</Text>
            <View style={styles.input}>
                <Feather name={iconsName} size={24} color="black" />
                <TextInput style={{
                    fontSize: 24,
                    width: '100%',
                    outlineColor: '#E2E2E2',
                }} placeholder={placeholder} placeholderTextColor={'#5D5F5F'}/>
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'column',
        // alignItems: 'center',
        gap: 8,
    },
    label: {
        fontSize: 40,
        fontWeight: 'semibold',
    },
    input: {
        padding: 0,
        flexDirection: 'row',
        gap: 8,
        borderRadius: 50,
        paddingLeft: 40,
        paddingVertical: 30,
        paddingRight: 10,
        backgroundColor: '#E2E2E2',
        color: 'black',
    }
})